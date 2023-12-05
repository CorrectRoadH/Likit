import { Button, Card, Drawer, Grid, Skeleton } from '@arco-design/web-react';
import React, { useState, useEffect } from 'react';
import styles from './style/overview.module.less';
import BusinessEditor from './business-editor';
import { BusinessType } from './type';

const { Row, Col } = Grid;
  
interface BusinessItemProps {
    business: BusinessType;
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
                  <div className={styles.title}>{business.title}</div>
                  <div className=''>{business.type}</div>
                </Skeleton>

                <Button type="text" size="small"
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
                    <BusinessEditor edit={true} />
                  </Drawer>
                </div>
              </Card>
            </Col>
    );
}
export default BusinessItem