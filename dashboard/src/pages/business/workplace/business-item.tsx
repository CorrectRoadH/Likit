import { Button, Card, Drawer, Grid, Skeleton, Typography } from '@arco-design/web-react';
import React, { useState } from 'react';
import styles from './style/overview.module.less';
import BusinessEditor from './business-editor';
import { Business } from '@/api/openapi';

const { Row, Col } = Grid;
  
const { Title,Text } = Typography;
interface BusinessItemProps {
    business: Business;
}

const BusinessItem = ({business}:BusinessItemProps) => {
    const [visible, setVisible] = useState(false);

    return(
        <Col span={6}>
          <Card className={styles['project-wrapper']} bordered={true} size="small">
              <div>
                <Skeleton 
                  loading={false} text={{ rows: 2, width: 60 }} animation
                >
                  <Title heading={5} className={styles.title}>{business.title}</Title>
                  <div>
                    <Text>Business ID</Text>
                    <Title heading={6}>{business.id}</Title>
                  </div>
                  
                  <div>{business.type}</div>
                </Skeleton>

                <Button size="small"
                  onClick={() => setVisible(true)}
                >
                  Configure
                </Button> 

                <Drawer  
                  width={600}
                  title={
                  <>
                      Edit Business
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
                    <BusinessEditor edit={true} 
                      business={business}
                    />
                  </Drawer>
                </div>
              </Card>
            </Col>
    );
}
export default BusinessItem