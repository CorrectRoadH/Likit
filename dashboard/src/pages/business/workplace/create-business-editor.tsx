import { Form, Input, Drawer, Button, Select } from "@arco-design/web-react";
import axios from "axios";
import React, { useState } from "react"

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
  


    return (
        <div>
      <Button
        onClick={() => {
          setVisible(true);
        }}
        type='primary'
      >
        Open Drawer
      </Button>
      <Drawer
        width={314}
        title={<span>Basic Information </span>}
        visible={visible}
        confirmLoading={confirmLoading}
        onOk={() => {
          form.validate().then((res) => {
            setConfirmLoading(true);

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
          {/* <Form.Item
            label='Self Introduction'
            required
            field='introduction'
            rules={[{ required: true }]}
          >
            <Input.TextArea placeholder='Plear enter' />
          </Form.Item> */}
        </Form>
      </Drawer>

        </div>
    )
}

export default CreateBusinessEditor