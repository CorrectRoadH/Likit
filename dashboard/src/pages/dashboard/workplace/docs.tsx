import React from 'react';
import { Link, Card, Typography } from '@arco-design/web-react';
import useLocale from '@/utils/useLocale';
import locale from './locale';
import styles from './style/docs.module.less';

const links = {
  github: 'https://github.com/CorrectRoadH/Likit',
  get_started: 'https://github.com/CorrectRoadH/Likit',
};
function QuickOperation() {
  const t = useLocale(locale);

  return (
    <Card>
      <div style={{ display: 'flex', justifyContent: 'space-between' }}>
        <Typography.Title heading={6}>{t['workplace.docs']}</Typography.Title>
        <Link href='https://github.com/CorrectRoadH/Likit'>{t['workplace.seeMore']}</Link>
      </div>
      <div className={styles.docs}>
        {Object.entries(links).map(([key, value]) => (
          <Link className={styles.link} key={key} href={value} target="_blank">
            {t[`workplace.${key}`]}
          </Link>
        ))}
      </div>
    </Card>
  );
}

export default QuickOperation;
