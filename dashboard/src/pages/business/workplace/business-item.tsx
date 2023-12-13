import { Button, Card, Drawer, Grid, Input, Skeleton, Typography } from '@arco-design/web-react';
import React, { useState } from 'react';
import styles from './style/overview.module.less';
import BusinessEditor from './business-editor';
import { Business } from '@/api/openapi';
import { useBusiness } from '@/api';

const { Row, Col } = Grid;
  
const { Title,Text } = Typography;
interface BusinessItemProps {
    business: Business;
}

const BusinessItem = ({business}:BusinessItemProps) => {
    const [visible, setVisible] = useState(false);
    const {updateBusiness} = useBusiness();

    const [title, setTitle] = useState(business.title);
    return(
        <Col span={6}>
          <Card className={styles['project-wrapper']} bordered={true} size="small">
              <div>
                <Skeleton 
                  loading={false} text={{ rows: 2, width: 60 }} animation
                >
                  <Title heading={5} className={styles.title}>{business.title}</Title>
                  {/* <Input value={title} onChange={(v)=>setTitle(v)} /> */}
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
                    console.log(business)
                    updateBusiness({
                      ...business,
                      title: title,
                    })
                    console.log("update")
                  }}
                  onCancel={() => setVisible(false)}
                >
                    <BusinessEditor edit={true}
                      title={title}
                      setTitle={setTitle} 
                      business={business}
                    />
                  </Drawer>
                </div>
              </Card>
            </Col>
    );
}
export default BusinessItem