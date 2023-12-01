import React from 'react';
import {
    Grid,
    Divider,
    Input,
    Select,
} from '@arco-design/web-react';

import styles from './style/overview.module.less';
import { BusinessType } from './type';

const { Row, Col } = Grid;

const Option = Select.Option;

interface BusinessEditorProps {
    business?: BusinessType;
    edit: boolean; // true: edit, false: create business
}

const BusinessEditor = ({edit}:BusinessEditorProps) => {
    return (
        <div>
            <h1>Business Editor</h1>

            <Row gutter={20}>
          <Col span={12}>
            <Input type='primary' placeholder="Business Title" />
          </Col>
          
          <Col span={12}>
            <Input placeholder="Business ID" />
          </Col>
        </Row>

        <Divider />

        <Row gutter={20}>
          <Col flex={1}>
            <Select placeholder="Business Type" defaultValue="1" >
                <Option value="1">Simple Vote System</Option>
                <Option value="2" disabled>Middle Vote System</Option>
            </Select>
          </Col>

          <Divider type="vertical" className={styles.divider}  />

          <Col flex={1}>
            <div>Qps:1000</div>
            <div>支持的功能:1000</div>
          </Col>
        </Row>
        </div>
    );
}
export default BusinessEditor