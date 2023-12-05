import { Form, Input, Drawer, Button, Select } from "@arco-design/web-react";
import axios from "axios";
import React, { useState,useEffect } from "react"
import { DatabaseConnectionConfig } from "../database/types";

const Option = Select.Option;

const formItemLayout = {
  wrapperCol: {
    span: 24,
  },
};

interface SystemFeature {
  features: string[],
  qps: number,
}

const SystemFeatures = {
  "SIMPLE_VOTE": {
    features: ["vote", "vote_count","list_voted_users"],
    qps: 100,
  },
}

const SystemFeatureTable = ({features,qps}:SystemFeature) => {
  return (
    <div>
      <div>
        <div>Support Features:</div>
        {
          features.map((item,index) => {
            return (
              <div key={index}>
                <span>{item}</span>
              </div>
            )
          })
        }
      </div>
      <span>qps:{qps}</span>
    </div>
  )
}


const CreateBusinessEditor = () => {

    const [visible, setVisible] = useState(false);
    const [form] = Form.useForm();
    const [confirmLoading, setConfirmLoading] = useState(false);

    const [data,setData] = useState<DatabaseConnectionConfig[]>([])

    const fetchData = () => {
      axios.get('/admin/v1/database').then((res) => {
          setData(res.data.dataSourceConfig);
      });
    }

    useEffect(() => {
        fetchData()
    }, []);

    return (
      <div>
        <Button
          onClick={() => {
            setVisible(true);
          }}
          type='primary'
        >
          Create Business
        </Button>
        <Drawer
          width={314}
          title={<span>Basic Information </span>}
          visible={visible}
          confirmLoading={confirmLoading}
          onOk={() => {
            form.validate().then((res) => {
              setConfirmLoading(true);
              res.config = {
                dataSourceConfig: []
              }

              // build the config for backend
              res.selectedDatabase.forEach((id) => {
                const database = data.find((item) => item.id === id)
                if (database) {
                  res.config.dataSourceConfig.push(database)
                }
              })

              console.log(res)
              setTimeout(() => {
                setVisible(false);
                setConfirmLoading(false);

                axios.post('/admin/v1/business', res).then((res) => {
                  console.log(res)
                })
              }, 1500);
            });
          }}
          onCancel={() => {
            setVisible(false);
          }}
        >
          <Form {...formItemLayout} form={form} layout='vertical'>
            <Form.Item label='Title' field='title' rules={[{ required: true }]}>
              <Input placeholder='Plear enter' />
            </Form.Item>
            <Form.Item label='Id' required field='id' rules={[{ required: true }]}>
              <Input placeholder='Plear enter' />
            </Form.Item>
            <Form.Item label='Vote System' field='type' rules={[{ required: true }]}>
              <Select placeholder='Plear select' options={['SIMPLE_VOTE']} />
            </Form.Item>

            <div>
              {
                form.getFieldValue('type') === 'SIMPLE_VOTE' 
                && 
                <SystemFeatureTable 
                  features={SystemFeatures[form.getFieldValue('type') ].features} 
                  qps={SystemFeatures[form.getFieldValue('type') ].qps} 
                />
              }
            </div>

            <Form.Item label='Database' field='selectedDatabase' rules={[{ required: true }]}>
              <Select placeholder='Please' mode='multiple'>
                {
                  data.map((item)=>(
                    <Option key={item.id} value={item.id}>
                      {item.title}
                    </Option>
                  ))
                }
              </Select> 
            </Form.Item>


          </Form>
        </Drawer>
      </div>
    )
}

export default CreateBusinessEditor