import React, { useState, useEffect } from 'react';
import {
  Grid,
  Card,
  Typography,
  Divider,
  Button,
  Input,
  Select,
  Skeleton,
} from '@arco-design/web-react';
import locale from './locale';
import useLocale from '@/utils/useLocale';
import axios from 'axios';
import styles from './style/overview.module.less';
const { Row, Col } = Grid;

const Option = Select.Option;

type DataType = {
  title?: string;
  business_id?: string;
  type?: string;
};


function Overview() {
  const [data, setData] = useState<DataType[]>([]);
  const [loading, setLoading] = useState(true);

  const t = useLocale(locale);

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


      <Row gutter={20}>
        <Col span={12}>
          <Input type='primary' placeholder="Business Title" />
        </Col>
        
        <Col span={12}>
          <Input placeholder="Business ID" />
        </Col>
      </Row>

      <Divider />

      <Row gutter={20}>
        <Col flex={1}>
          <Select placeholder="Business Type" defaultValue="1" >
              <Option value="1">Simple Vote System</Option>
              <Option value="2" disabled>Middle Vote System</Option>
          </Select>
        </Col>

        <Divider type="vertical" className={styles.divider}  />

        <Col flex={1}>
          <div>Qps:1000</div>
          <div>支持的功能:1000</div>
        </Col>
      </Row>

      <Divider />

      <Row>
        <Button type="primary" style={{ marginLeft: 8 }}>新建 Business</Button>
      </Row>

      <Divider />
      
      全部的 Business

      <Row gutter={20}>
        {
          data.map((item)=>
            <Col span={4} key={item.business_id} >
              <div className={styles.item}>
                <div></div>
                <div>
                  <Skeleton 
                    loading={loading} text={{ rows: 2, width: 60 }} animation
                  >
                    <div className={styles.title}>{item.title}</div>
                    <div className=''>{item.type}</div>
                  </Skeleton>

                  <Button type="text" size="small">
                    Configure
                  </Button> 

                </div>
              </div>
            </Col>
          )
        }
      </Row>

      <Divider />
    </Card>
  );
}

export default Overview;
