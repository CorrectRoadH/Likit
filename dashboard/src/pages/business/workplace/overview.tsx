import React, { useState, useEffect, ReactNode } from 'react';
import {
  Grid,
  Card,
  Typography,
  Divider,
  Button,
  Input,
  Select,
} from '@arco-design/web-react';
import locale from './locale';
import useLocale from '@/utils/useLocale';

const { Row, Col } = Grid;

const Option = Select.Option;


function Overview() {
  const t = useLocale(locale);


  return (
    <Card>
      <Typography.Title heading={5}>
        {'Business Overview'}
      </Typography.Title>
      <Divider />

        <Col flex={1}>
           
            <Input placeholder="Business Title" />
            
            <Row>
                <Select placeholder="Business Type" >
                    <Option value="1">Simple Vote System</Option>
                    <Option value="2" disabled>Middle Vote System</Option>
                </Select>
                
                <div>Qps:1000</div>
                <div>支持的功能:1000</div>
            </Row>
        </Col>

        <Button type="primary" style={{ marginLeft: 8 }}>新建 Business</Button>
      <Divider />
        全部的 Business

      <Divider />
    </Card>
  );
}

export default Overview;
