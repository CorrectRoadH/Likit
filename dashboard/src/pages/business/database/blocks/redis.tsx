import React from 'react'
import { DatabaseConnectionConfig } from '../../../../types'
import styles from '../style/blocks.module.less';
import { Button, Card,Typography } from '@arco-design/web-react';
const { Text, Title } = Typography;

interface RedisProps {
    config: DatabaseConnectionConfig
}
const Redis = ({config}:RedisProps)=>{
    return (
        <Card className={styles['project-wrapper']} bordered={true} size="small">
            <Title heading={5}>{config.title}</Title>


            <Text>{`${config.host}:${config.port}`}</Text>
            <div>
                <Button>Configure</Button>
            </div>
        </Card>
    )
}

export default Redis