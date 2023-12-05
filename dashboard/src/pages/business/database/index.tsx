import React from 'react';
import { Grid, Space } from '@arco-design/web-react';
import Overview from './overview';
import { Toaster } from 'sonner'
import styles from './style/index.module.less';

const { Row, Col } = Grid;

function Workplace() {
    return (
        <div className={styles.wrapper}>
            <Toaster />
            <Space size={16} direction="vertical" className={styles.left}>
                <Overview />
                <Row gutter={16}>
                    <Col span={12}>
                    </Col>
                    <Col span={12}>
                    </Col>
                </Row>
            </Space>

        </div>
    )
}

export default Workplace;