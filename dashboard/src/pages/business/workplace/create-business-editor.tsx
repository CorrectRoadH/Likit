import { Form, Input, Drawer, Button, Select, Tag, Space, Typography, Skeleton, Popover } from "@arco-design/web-react";
import React, { useState } from "react"
import { DatabaseConnectionConfig } from "../../../types";
import { toast } from "sonner";
import style from './style/overview.module.less'
import useLocale from "@/utils/useLocale";
import locale from "./locale";
import { useBusiness, useDatabase } from "@/api";
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
    id: "SIMPLE_VOTE",
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

const system_features = ['SIMPLE_VOTE']

const VoteSystemComponent = ({features,qps,require}:SystemFeature) => {
  const t = useLocale(locale);

  return (
    <div>
      <div>
      <Title heading={6}>Features:</Title>
        <Space className={style.features}>
          {
            features.map((item,index) => 
                <Popover
                  key={index}
                  title={t[`vote.description.${item}`]}
                >
                  <Tag  >
                    {t[`vote.${item}`]}
                  </Tag>
                </Popover>
              )
          }
        </Space>
      </div>

      <div>
        <Title heading={6}>Qbs:</Title>
        <Title  heading={2}>
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
                  <Title heading={2}>{item.number}</Title>
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

    // const [data,setData] = useState<DatabaseConnectionConfig[]>([])

    const [beSelectedSystem,setBeSelectedSystem] = useState<string>()
    const t = useLocale(locale);

    const { createBusiness } = useBusiness()
    const { database, isLoading, isError } = useDatabase()

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
          width={600}
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
                const db = database.find((item) => item.id === id)
                if (db) {
                  res.config.dataSourceConfig.push(db)
                }
              })

              console.log(res)
              setTimeout(() => {
                setVisible(false);
                setConfirmLoading(false);

                createBusiness(res).then((res) => {
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

            <Form.Item label='Vote System' field='type' rules={[{ required: true }]}
            >
              <Select placeholder='System select'
                onChange={(value) => {
                  setBeSelectedSystem(value)
                }}
              >
                {
                  system_features.map((item) => {
                    return (
                      <Option key={item} value={item}>
                       {t[`vote.system.${item}`]}
                      </Option>
                    )
                  })
                }
              </Select>
            </Form.Item>

            <div>
              {
                form.getFieldValue('type') === 'SIMPLE_VOTE' 
                && 
                <VoteSystemComponent 
                  features={SystemFeatures[form.getFieldValue('type') ].features} 
                  qps={SystemFeatures[form.getFieldValue('type') ].qps} 
                  require={SystemFeatures[form.getFieldValue('type') ].require}
                />
              }
            </div>

            <Skeleton
              loading={isLoading}
            >
              <Form.Item label='Database' field='selectedDatabase' rules={[{ required: true }]}>
                <Select placeholder='Please' mode='multiple'>
                  {
                    !isLoading && database.map((item)=>(
                      <Option key={item.id} value={item.id}>
                        {item.title}
                      </Option>
                    ))
                  }
                </Select> 
              </Form.Item>
            </Skeleton>

          </Form>
        </Drawer>
      </div>
    )
}

export default CreateBusinessEditor