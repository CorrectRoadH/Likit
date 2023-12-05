import React, { useState, useEffect } from 'react';
import {
  Grid,
  Card,
  Typography,
  Divider,
  Button,
  Skeleton,
  Drawer,
} from '@arco-design/web-react';
import locale from './locale';
import useLocale from '@/utils/useLocale';
import axios from 'axios';
import styles from './style/overview.module.less';
import BusinessEditor from './business-editor';
import BusinessItem from './business-item';
import { BusinessType } from './type';
import CreateBusinessEditor from './create-business-editor';

const { Row, Col } = Grid;


function Overview() {
  const [data, setData] = useState<BusinessType[]>([]);
  const [loading, setLoading] = useState(true);

  const t = useLocale(locale);

  const [visible, setVisible] = useState(false);

  const fetchData = () => {
    setLoading(true);
    axios
      .get('/admin/v1/businesses')
      .then((res) => {
        console.log(res)
        setData(res.data);
      })
      .finally(() => {
        setLoading(false);
      });
  };

  useEffect(() => {
    fetchData();
  }, []);

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
          data.map((item)=>
            <BusinessItem key={item.business_id} business={item} />
          )
        }
      </Row>

      <Divider />
    </Card>
  );
}

export default Overview;
