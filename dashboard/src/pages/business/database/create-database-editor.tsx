import { useDatabase } from "@/api";
import { DatabaseConnectConfig } from "@/api/openapi";
import { uuid } from "@/utils/uuid";
import { Form, Button, Drawer, Input, Select, Skeleton } from "@arco-design/web-react"
import axios from "axios";
import React, { useState } from "react"
import { toast } from "sonner";

const formItemLayout = {
    wrapperCol: {
      span: 24,
    },
  };

const database_require = {
    "redis":{
        "host":[{ required: true }],
        "username":[],
        "password":[],
        "port":[{ required: true }],
        "database":[],
    },
    "postgres":{
        "host":[{ required: true }],
        "username":[{ required: true }],
        "password":[{ required: true }],
        "port":[{ required: true }],
        "database":[{ required: true }],
    }
}
  
  
const CreateDatabase = () => {
    const [visible, setVisible] = useState(false);
    const [form] = Form.useForm();
    const [confirmLoading, setConfirmLoading] = useState(false);
    const { createDatabase } = useDatabase()

    const [ beSelectedType, setBeSelectedType ] = useState("")

    const handleTestBtnClick = () => {
        form.validate().then((res)=>{
            const config:DatabaseConnectConfig = {
                title: res.title,
                "id": uuid(),
                "databaseType":"redis",
                "host": res.host,
                "port": Number(res.port),
                "username":"",
                password: res?.password || "",
                "database":""
            }
            console.log(config)
            axios.post("/admin/v1/database/test", config).then((res) => {
                toast.success("connect success");
            }).catch((err) => {
                toast.error(err.response.data.msg);
            })    
        })
    }


    return <>
        <Button type="primary"
            onClick={() => {
                setVisible(true);
            }}
        >Create Database</Button>
        <Drawer
            width={600}
            title={<span>Basic Information </span>}
            visible={visible}
            confirmLoading={confirmLoading}

            okText={"Create"}
            cancelText={"Cancel"}
            onOk={()=>{
                form.validate().then((res)=>{
                    setConfirmLoading(true);
                    console.log(res)
                    createDatabase(res).then((res)=>{
                        toast.success("create success");
                    })
                    setConfirmLoading(false);
                })
            }}

            onCancel={() => {
                setVisible(false);
              }}    
        >

            <Form
                {...formItemLayout}
                form={form}
                layout="vertical"
            >
                <Form.Item label='Database' field='type' rules={[{ required: true }]}>
                    <Select placeholder='Please select a database type' 
                        options={['redis','postgres']} 
                        onChange={(value)=>{
                            setBeSelectedType(value)
                        }}
                    />
                </Form.Item>
                
                {
                beSelectedType != "" && 
                <Skeleton
                    loading={false}
                >
                    <Form.Item label='Title' field='title' rules={[{ required: true }]}>
                        <Input placeholder='Title' />
                    </Form.Item>

                    <Form.Item label='Host' field='host' 
                        rules={database_require[beSelectedType]['host']}
                    >
                        <Input placeholder='Host' />
                    </Form.Item>

                    <Form.Item label='Port' field='port' 
                        rules={database_require[beSelectedType]['port']}
                    >
                        <Input placeholder='Port' />
                    </Form.Item>

                    <Form.Item label='Username' field='username'
                        rules={database_require[beSelectedType]['username']}
                    >
                        <Input placeholder='Username' />
                    </Form.Item>

                    <Form.Item label='Password' field='password' 
                        rules={database_require[beSelectedType]['password']}
                    >
                        <Input placeholder='Password' />
                    </Form.Item>

                    <Form.Item label='Database' field='database'
                        rules={database_require[beSelectedType]['database']}
                    >
                        <Input placeholder='Input the Database' />
                    </Form.Item>

                    <Button
                        onClick={handleTestBtnClick}
                    >test connect</Button>
                </Skeleton>
            }
            </Form>
        </Drawer>
    </>
}
export default CreateDatabase