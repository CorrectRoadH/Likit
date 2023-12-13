import React, { useState } from 'react'
import { DatabaseConnectConfig } from "@/api/openapi";
import styles from '../style/blocks.module.less';
import { Button, Card,Drawer,Typography } from '@arco-design/web-react';
import RedisEditor from '../editor/redis-editor';
const { Text, Title } = Typography;

interface PostgresProps {
    config: DatabaseConnectConfig
}
const Postgres = ({config}:PostgresProps)=>{
    const [visible, setVisible] = useState(false);

    return (
        <Card className={styles['project-wrapper']} bordered={true} size="small">
            <Title heading={5}>{config.title}</Title>

            <Text>{`${config.host}:${config.port}`}</Text>

            <div>
                <Button
                    onClick={() => setVisible(true)}
                >Configure</Button>
                <Drawer
                    width={600}
                    title={
                    <>
                        Edit Redis
                    </>
                    }
                    visible={visible}
                    okText={"Update"}
                    cancelText={"Cancel"}
                    onOk={()=>{
                      console.log("update")
                    }}
                    onCancel={() => setVisible(false)}  
                >
                    <RedisEditor database={config} />
                </Drawer>
            </div>

        </Card>
    )
}

export default Postgres