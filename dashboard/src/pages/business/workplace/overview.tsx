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

const { Row, Col } = Grid;


function Overview() {
  const [data, setData] = useState<BusinessType[]>([]);
  const [loading, setLoading] = useState(true);

  const t = useLocale(locale);

  const [visible, setVisible] = useState(false);


  const handleNewBusiness = () => {
    console.log("hello")
  }
  const fetchData = () => {
    setLoading(true);
    axios
      .get('/api/v1/businesses')
      .then((res) => {
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

      <Row>
        <Button type="primary" style={{ marginLeft: 8 }}
          onClick={() => setVisible(true)}
        >Create Business</Button>
      </Row>

      <Drawer
        width={600}
        title={
          <>
            New Business
          </>
        }
        visible={visible}
        okText={"Create"}
        cancelText={locale['settings.close']}
        onOk={handleNewBusiness}
        onCancel={() => setVisible(false)}
      >
        <BusinessEditor edit={false} />
      </Drawer>

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
