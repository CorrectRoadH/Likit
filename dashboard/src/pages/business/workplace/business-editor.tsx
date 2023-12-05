import React, { useEffect, useState } from 'react';
import {
    Grid,
    Divider,
    Input,
    Select,
    Skeleton,
} from '@arco-design/web-react';

import styles from './style/overview.module.less';
import { BusinessType } from './type';
import axios from 'axios';

const { Row, Col } = Grid;

const Option = Select.Option;

interface BusinessEditorProps {
    business?: BusinessType;
    edit: boolean; // true: edit, false: create business
}

interface VoteSystem {
  id: string,
  feature: string[],
  qps: number,
}

const BusinessEditor = ({business}:BusinessEditorProps) => {
  const [data, setData] = useState<VoteSystem[]>([]);
  const [loading, setLoading] = useState(true);
 
  const fetchData = () => {
    setLoading(true);
    axios
      .get('/admin/v1/vote_system')
      .then((res) => {
        console.log(data)
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
        <div>
          <h1>Business Editor</h1>

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
              <Skeleton
                loading={loading}
                text={{rows:2, width:60}}
                animation
              >
              </Skeleton>
            </Col>
          </Row>
        </div>
    );
}
export default BusinessEditor