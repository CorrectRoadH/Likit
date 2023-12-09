import { Form, Input, Drawer, Button, Select, Tag, Space, Typography } from "@arco-design/web-react";
import axios from "axios";
import React, { useState,useEffect } from "react"
import { DatabaseConnectionConfig } from "../database/types";
import { toast } from "sonner";
import style from './style/overview.module.less'
const Option = Select.Option;

const { Title } = Typography

const formItemLayout = {
  wrapperCol: {
    span: 24,
  },
};

interface SystemFeature {
  features: string[],
  qps: number,
  require: {
    type: string,
    number: number,
  }[],
}

const SystemFeatures = {
  "SIMPLE_VOTE": {
    features: ["vote","unvote", "vote_count","list_voted_users","is_voted"],
    qps: 100,
    require: [
      {
        type: "redis",
        number: 1,
      }
    ],
  },
}

const SystemFeatureTable = ({features,qps,require}:SystemFeature) => {
  return (
    <div>
      <div>
      <Title heading={6}>Features:</Title>
        <Space className={style.features}>
          {
            features.map((item,index) => {
              return (
                <Tag  key={index}>
                  {item}
                </Tag>
              )
            })
          }
        </Space>
      </div>

      <div>
        <Title heading={6}>Qbs:</Title>
        <Title>
          {qps}
        </Title>
      </div>

      <div>
        <Title heading={6}>Require:</Title>
        <Space>
          {
            require.map((item,index) => {
              return (
                <div key={index}>
                  <span>{item.type}</span>
                  <Title>{item.number}</Title>
                </div>
              )
            })
          }
        </Space>
      </div>
    </div>
  )
}


const CreateBusinessEditor = () => {

    const [visible, setVisible] = useState(false);
    const [form] = Form.useForm();
    const [confirmLoading, setConfirmLoading] = useState(false);

    const [data,setData] = useState<DatabaseConnectionConfig[]>([])

    const [beSelectedSystem,setBeSelectedSystem] = useState<string>()

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
                  toast.success('Create Business Success')
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
              <Input placeholder='Business title' />
            </Form.Item>

            <Form.Item label='Id' required field='id' rules={[{ required: true }]}>
              <Input placeholder='Business ID. like: COMMENT_LIKE' />
            </Form.Item>


            {/* TODO rerender the component when select changes */}
            <Form.Item label='Vote System' field='type' rules={[{ required: true }]}
            >
              <Select placeholder='System select' options={['SIMPLE_VOTE']} 
                onChange={(value) => {
                  setBeSelectedSystem(value)
                }}
              />
            </Form.Item>

            <div>
              {
                form.getFieldValue('type') === 'SIMPLE_VOTE' 
                && 
                <SystemFeatureTable 
                  features={SystemFeatures[form.getFieldValue('type') ].features} 
                  qps={SystemFeatures[form.getFieldValue('type') ].qps} 
                  require={SystemFeatures[form.getFieldValue('type') ].require}
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