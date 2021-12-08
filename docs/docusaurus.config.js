const lightCodeTheme = require('prism-react-renderer/themes/github');
const darkCodeTheme = require('prism-react-renderer/themes/dracula');

/** @type {import('@docusaurus/types').DocusaurusConfig} */
module.exports = {
  title: 'go-cf-api',
  tagline: 'Golang Implementation of CloudController_NG',
  url: 'https://pages.github.tools.sap/',
  baseUrl: '/cloudfoundry/go-cf-api/',
  onBrokenLinks: 'throw',
  onBrokenMarkdownLinks: 'warn',
  favicon: 'img/favicon.ico',
  organizationName: 'cloudfoundry',
  projectName: 'go-cf-api',
  i18n: {
    defaultLocale: 'en',
    locales: ['en'],
  },
  themeConfig: {
    navbar: {
      title: 'go-cf-api',
      logo: {
        alt: 'Documentation of the go-cf-api Project',
        src: 'img/logo.svg',
      },
      items: [
        {
          type: 'docsVersionDropdown',
          position: 'right',
          dropdownItemsAfter: [{to: '/versions', label: 'All versions'}],
          dropdownActiveClassDisabled: true,
        },
        {
          type: 'localeDropdown',
          position: 'right',
          dropdownItemsAfter: [
            {
              to: '/help-us-translate',
              label: 'Help us translate',
            },
          ],
        },
        {
          to: '/docs/intro',
          label: 'Documentation',
          position: 'left',
          activeBaseRegex: `/docs/`,
        },
        {
          to: '/godocs/intro',
          label: 'GoDocs',
          position: 'left',
          activeBaseRegex: `/godocs/`,
        },
        {to: '/api', label: 'API Docs', position: 'left'},
        {to: '/adrs', label: 'Architectural Decision Records', position: 'left'},
        {
          href: 'https://github.tools.sap/cloudfoundry/go-cf-api',
          label: 'GitHub',
          position: 'right',
        },
      ],
    },
    footer: {
      style: 'dark',
      links: [
        {
          title: 'Documentation',
          items: [
            {
              label: 'API',
              to: '/api',
            },
            {
              label: 'Documentation',
              to: '/docs/intro',
            },
            {
              label: 'GoLang Documentation',
              to: '/godocs/intro',
            },
          ],
        },
        {
          title: 'More',
          items: [
            {
              label: 'Architectural Decision Records',
              to: '/adrs',
            },
            {
              label: 'GitHub',
              href: 'https://github.tools.sap/cloudfoundry/go-cf-api',
            },
          ],
        },
      ],
      copyright: `Copyright © ${new Date().getFullYear()} SAP SE`,
    },
    prism: {
      theme: lightCodeTheme,
      darkTheme: darkCodeTheme,
    },
  },
  presets: [
    [
      '@docusaurus/preset-classic',
      {
        docs: {
          id: 'docs',
          path: 'docs',
          editUrl: ({versionDocsDirPath, docPath}) =>
              `https://github.com/cloudfoundry/go-cf-api/edit/main/docs/${versionDocsDirPath}/${docPath}`,
          editLocalizedFiles: true,
          editCurrentVersion: true,
          routeBasePath: 'docs',
          include: ['**/*.md', '**/*.mdx'],
          exclude: [
            '**/_*.{js,jsx,ts,tsx,md,mdx}',
            '**/_*/**',
            '**/*.test.{js,jsx,ts,tsx}',
            '**/__tests__/**',
            '**/Readme.md',
            '**/README.md'
          ],
          sidebarPath: require.resolve('./sidebars.js'),
          docLayoutComponent: '@theme/DocPage',
          docItemComponent: '@theme/DocItem',
          showLastUpdateAuthor: true,
          showLastUpdateTime: true,
          disableVersioning: false,
          includeCurrentVersion: true,
          lastVersion: undefined,
          versions: {
            current: {
              label: 'Latest',
            },
          },
        },
        blog: {
          path: 'adrs',
          editUrl: ({locale, blogDirPath, blogPath, permalink}) => {
            return `https://github.com/cloudfoundry/go-cf-api/edit/main/docs/${blogDirPath}/${blogPath}`;
          },
          editLocalizedFiles: true,
          blogTitle: 'Architectural Decision Records',
          blogDescription: 'ADRs',
          blogSidebarCount: 5,
          blogSidebarTitle: 'All our posts',
          routeBasePath: 'adrs',
          include: ['**/*.{md,mdx}'],
          exclude: [
            '**/_*.{js,jsx,ts,tsx,md,mdx}',
            '**/_*/**',
            '**/*.test.{js,jsx,ts,tsx}',
            '**/__tests__/**',
            '**/Readme.md',
            '**/README.md'
          ],
          postsPerPage: 10,
          blogListComponent: '@theme/BlogListPage',
          blogPostComponent: '@theme/BlogPostPage',
          blogTagsListComponent: '@theme/BlogTagsListPage',
          blogTagsPostsComponent: '@theme/BlogTagsPostsPage',
          showReadingTime: true,
        },
        theme: {
          customCss: require.resolve('./src/css/custom.css'),
        },
      },
    ],
    [
      'redocusaurus',
      {
        specs: [{
          spec: 'swagger.yaml',
          routePath: '/api/',
        }],
        theme: {
          primaryColor: '#00457f',
          redocOptions: { hideDownloadButton: false },
        },
      }
    ],
  ],
  plugins:[
    [
      '@docusaurus/plugin-content-docs',
      {
        path: 'godocs',
        editUrl: ({versionDocsDirPath, docPath}) =>
            `https://github.com/cloudfoundry/go-cf-api/edit/main/docs/${versionDocsDirPath}/${docPath}`,
        editLocalizedFiles: true,
        editCurrentVersion: true,
        routeBasePath: 'godocs',
        include: ['**/*.md', '**/*.mdx'],
        exclude: [
          '**/_*.{js,jsx,ts,tsx,md,mdx}',
          '**/_*/**',
          '**/*.test.{js,jsx,ts,tsx}',
          '**/__tests__/**',
          '**/Readme.md',
          '**/README.md'
        ],
        sidebarPath: require.resolve('./sidebars.js'),
        docLayoutComponent: '@theme/DocPage',
        docItemComponent: '@theme/DocItem',
        remarkPlugins: [],
        rehypePlugins: [],
        beforeDefaultRemarkPlugins: [],
        beforeDefaultRehypePlugins: [],
        showLastUpdateAuthor: true,
        showLastUpdateTime: true,
        disableVersioning: false,
        includeCurrentVersion: true,
        lastVersion: undefined,
        versions: {
          current: {
            label: 'Latest',
          },
        },
      },
    ]
  ],
};
