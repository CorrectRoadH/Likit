import React from 'react';
import {
  Grid,
  Card,
  Typography,
  Divider,
} from '@arco-design/web-react';
import styles from './style/overview.module.less';
import BusinessItem from './business-item';
import { BusinessType } from '@/types';
import CreateBusinessEditor from './create-business-editor';
import { useBusiness } from '@/api';

const { Row } = Grid;


function Overview() {
  const { businesses, isLoading, isError} = useBusiness()

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

      <CreateBusinessEditor />


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
