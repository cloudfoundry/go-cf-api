package metadata

import (
	"fmt"
	"regexp"
	"strings"
)

//go:generate mockgen -source=$GOFILE -destination=mocks/$GOFILE
type LabelSelectorParser interface {
	Parse(labelSelector string) (LabelSelectorFilters, error)
}

const (
	Exists    Operator = ""
	NotExists Operator = "!"
	Eq        Operator = "="
	EqAlt     Operator = "=="
	NotEq     Operator = "!="
	In        Operator = "in"
	NotIn     Operator = "notin"

	AllowedChars          = "[[:alnum:]][[:alnum:]._-]*"  // allow chars in keys: a-z,A-z,0-9,-,_,.
	AllowedCharsWithComma = "[[:alnum:]][[:alnum:]._,-]*" // in clauses can contain commas

)

type (
	Key      string
	Operator string

	LabelSelector struct {
		Key      Key
		Operator Operator
		Values   []string
	}
	LabelSelectorSlice  []LabelSelector
	labelSelectorParser struct{}
)

func NewLabelSelectorParser() LabelSelectorParser {
	return &labelSelectorParser{}
}

func (l *labelSelectorParser) Parse(labelSelector string) (LabelSelectorFilters, error) {
	// Split on , unless it is within a set of brackets (in clause)
	selector := regexp.MustCompile(`(?:\(.*?\)|[^,])+`)
	//nolint:varnamelen
	in := regexp.MustCompile(fmt.Sprintf(`^(%s) ((?:not)?in) \((%s)\)$`, AllowedChars, AllowedCharsWithComma))
	equal := regexp.MustCompile(fmt.Sprintf(`^(%s)([!=]?=)(%s)$`, AllowedChars, AllowedChars))
	exists := regexp.MustCompile(fmt.Sprintf(`^(!?)(%s)$`, AllowedChars))

	var labelSelectorSlice LabelSelectorSlice
	//nolint:nestif  //Regex is ugly
	for _, elem := range selector.FindAllString(labelSelector, -1) {
		if match := in.FindAllStringSubmatch(elem, 1); match != nil {
			if err := labelSelectorSlice.append(Key(match[0][1]), Operator(match[0][2]), strings.Split(match[0][3], ",")); err != nil {
				return nil, err
			}
		} else if match := equal.FindAllStringSubmatch(elem, 1); match != nil {
			if err := labelSelectorSlice.append(Key(match[0][1]), Operator(match[0][2]), []string{match[0][3]}); err != nil {
				return nil, err
			}
		} else if match := exists.FindAllStringSubmatch(elem, 1); match != nil {
			var operator Operator
			if match[0][1] != string(NotExists) {
				operator = Exists
			} else {
				operator = NotExists
			}
			if err := labelSelectorSlice.append(Key(match[0][2]), operator, nil); err != nil {
				return nil, err
			}
		} else {
			return nil, fmt.Errorf("could not parse label selectors '%s'", labelSelector)
		}
	}
	return labelSelectorSlice, nil
}

func (k Key) isValid() bool {
	var firstLastRegex *regexp.Regexp
	switch {
	case len(k) > 1:
		firstLastRegex = regexp.MustCompile(fmt.Sprintf(`^(%s).*(%s)\z`, AllowedChars, AllowedChars))
	case len(k) == 1:
		firstLastRegex = regexp.MustCompile(fmt.Sprintf(`^(%s)\z`, AllowedChars))
	default:
		return false
	}

	return firstLastRegex.MatchString(string(k))
}

func (op Operator) isValid() bool {
	const t = true
	valid := map[Operator]bool{Exists: t, NotExists: t, Eq: t, EqAlt: t, NotEq: t, In: t, NotIn: t}
	return valid[op]
}

func (selectors *LabelSelectorSlice) append(key Key, operator Operator, values []string) error {
	if !key.isValid() {
		return fmt.Errorf("%s starts or ends with invalid characters", key)
	}

	if !operator.isValid() {
		return fmt.Errorf("invalid label_selector value")
	}

	if map[Operator]bool{Exists: true, NotExists: true}[operator] {
		if len(values) != 0 {
			return fmt.Errorf("invalid label_selector value")
		}
	}

	if map[Operator]bool{Eq: true, EqAlt: true, NotEq: true}[operator] {
		if len(values) != 1 {
			return fmt.Errorf("invalid label_selector value")
		}
	}

	*selectors = append(*selectors, LabelSelector{
		Key:      key,
		Operator: operator,
		Values:   values,
	})
	return nil
}
