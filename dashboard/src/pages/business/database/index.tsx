import React from 'react';
import { Grid, Space } from '@arco-design/web-react';
import Overview from './overview';
import { Toaster } from 'sonner'

const { Row, Col } = Grid;

function Workplace() {
    return (
        <div >
            <Toaster />
            <Space size={16} direction="vertical">
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