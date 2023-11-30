import React from 'react';
import styles from './style/index.module.less';
import { Grid, Space } from '@arco-design/web-react';
import Overview from './overview';

const { Row, Col } = Grid;

const gutter = 16;

function Workplace() {
    return (
        <div className={styles.wrapper}>

            <Space size={16} direction="vertical" className={styles.left}>
                <Overview />
                <Row gutter={gutter}>
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