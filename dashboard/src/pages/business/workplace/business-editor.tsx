import React, { useEffect, useState } from 'react';
import {
    Grid,
    Divider,
    Input,
    Select,
    Skeleton,
    Button,
    Modal,
} from '@arco-design/web-react';

import styles from './style/overview.module.less';
import { BusinessType } from '../../../types/type';
import axios from 'axios';
import { toast } from 'sonner';

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

  const handleDeleteBtnClick = ()=>{
    Modal.confirm({
      title: 'Are you sure to delete this business?',
      content: 'This action cannot be undone.',
      onOk: () => {
        axios.delete(`/admin/v1/business?id=${business.id}`).then((res)=>{
          toast.success('Delete business successfully!')
          // refresh the page
          window.location.reload();
        })
      },
    });
  }

  return (
        <div>
          <h1>Business Editor</h1>


          <Row gutter={20}>
            <Col span={12}>
              <Input type='primary' placeholder="Business Title"
                value={business.title}
              />
            </Col>
            
            <Col span={12}>
              <Input placeholder="Business ID" 
                disabled
                value={business.id}
              />
            </Col>
          </Row>

          <Divider />

          {/* <Row gutter={20}>
            <Col flex={1}>
              <Skeleton
                loading={loading}
                text={{rows:2, width:60}}
                animation
              >
              </Skeleton>
            </Col>
          </Row> */}

          <Row>
            <Col span={12}>
              <Button
                onClick={handleDeleteBtnClick}
              >Delete</Button>
            </Col>
          </Row>
        </div>
    );
}
export default BusinessEditor