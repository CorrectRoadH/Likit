import React, { useState, useEffect } from 'react';
import {
  Grid,
  Card,
  Typography,
  Divider,
} from '@arco-design/web-react';
import locale from './locale';
import useLocale from '@/utils/useLocale';
import axios from 'axios';
import styles from './style/overview.module.less';
import BusinessItem from './business-item';
import { BusinessType } from './type';
import CreateBusinessEditor from './create-business-editor';
import { useBusiness } from '@/api';

const { Row, Col } = Grid;


function Overview() {
  const { businesses,isLoading,isError} = useBusiness()

  const t = useLocale(locale);

  return (
    <Card>
      <Typography.Title heading={5}>
        {'Business Overview'}
      </Typography.Title>
      <Divider />

      <div className={styles.ctw}>
        <Typography.Paragraph
          className={styles['chart-title']}
          style={{ marginBottom: 0 }}
        >
          Business Workspace
        </Typography.Paragraph>
      </div>

      {/* <Row> */}
        <CreateBusinessEditor />
      {/* </Row> */}


      <Divider />
      
      
      <div className={styles.ctw}>
        <Typography.Paragraph
          className={styles['chart-title']}
          style={{ marginBottom: 0 }}
        >
          Business Overview
        </Typography.Paragraph>
      </div>

      <Row gutter={20}>
        {
          businesses.map((item)=>
            <BusinessItem key={item.id} business={item} />
          )
        }
      </Row>

    </Card>
  );
}

export default Overview;
