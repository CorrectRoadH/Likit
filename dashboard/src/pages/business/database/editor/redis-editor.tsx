import { useDatabase } from '@/api';
import { DatabaseConnectConfig } from '@/api/openapi';
import {
    Grid,
    Divider,
    Input,
    Button,
    Modal,
} from '@arco-design/web-react';
import React from "react";
import { toast } from 'sonner'

const { Row, Col } = Grid;

interface RedisEditorProps {
    database?: DatabaseConnectConfig;
}


const RedisEditor = ({database}:RedisEditorProps) => {


  const {deleteDatabase} = useDatabase();
  const handleDeleteBtnClick = ()=>{
    Modal.confirm({
      title: 'Are you sure to delete this database?',
      content: 'This action cannot be undone.',
      onOk: () => {
        deleteDatabase(database.id).then((res)=>{
          toast.success('Delete business successfully!')
        })
      },
    });
  }


    return (
        <div>
          <h1>Database Editor</h1>

          <Row gutter={20}>
            <Col span={12}>
              <Input type='primary' placeholder="Business Title"
                value={database.title}
              />
            </Col>
            
            <Col span={12}>
              <Input placeholder="Business ID" 
                disabled
                value={database.id}
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
    )
}

export default RedisEditor;