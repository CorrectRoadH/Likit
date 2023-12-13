import React, { useEffect, useState } from 'react';
import {
    Grid,
    Divider,
    Input,
    Button,
    Modal,
} from '@arco-design/web-react';

import axios from 'axios';
import { toast } from 'sonner';
import { Business } from '@/api/openapi';
import { useBusiness } from '@/api';

const { Row, Col } = Grid;


interface BusinessEditorProps {
    business?: Business;
    edit: boolean; // true: edit, false: create business
    title: string;
    setTitle: (title: string) => void;
}

interface VoteSystem {
  id: string,
  feature: string[],
  qps: number,
}

const BusinessEditor = ({business,title ,setTitle}:BusinessEditorProps) => {
  const [data, setData] = useState<VoteSystem[]>([]);
  const [loading, setLoading] = useState(true);
 
  const {deleteBusiness} = useBusiness();

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
        deleteBusiness(business.id).then((res)=>{
          toast.success('Delete business successfully!')
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
                value={title}
                onChange={setTitle}
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