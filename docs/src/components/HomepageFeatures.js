import React from 'react';
import clsx from 'clsx';
import styles from './HomepageFeatures.module.css';

const FeatureList = [
  {
    title: 'CloudGontroller Documentation',
    Svg: require('../../static/img/logo.svg').default,
    description: (
      <>
          <a href="/cloudfoundry/cloudgontroller/docs/intro">Documentation</a> - General documentation on usage and development<br/>
          <a href="/cloudfoundry/cloudgontroller/godocs/intro">Golang Documentation</a> - Generated package documentation from sourcecode comments<br/>
          <a href="/cloudfoundry/cloudgontroller/api">API Documentration</a> - What is implemented so far from <a href="https://v3-apidocs.cloudfoundry.org/">Cloud Controller V3 API</a><br/>
          <a href="/cloudfoundry/cloudgontroller/adrs">Achitectural Decition Records</a> - Design Documents for this project<br/>
      </>
    ),
  },
];

function Feature({Svg, title, description}) {
  return (
    <div className={clsx('col col--20')}>
      {/*<div className="text--center">*/}
      {/*  <Svg className={styles.featureSvg} alt={title} />*/}
      {/*</div>*/}
      <div className="text--center padding-horiz--md">
        <h3>{title}</h3>
        <p>{description}</p>
      </div>
    </div>
  );
}

export default function HomepageFeatures() {
  return (
    <section className={styles.features}>
      <div className="container">
        <div className="row">
          {FeatureList.map((props, idx) => (
            <Feature key={idx} {...props} />
          ))}
        </div>
      </div>
    </section>
  );
}