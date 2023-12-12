import React from "react";
import { Card, Divider, Typography,Grid, Skeleton } from "@arco-design/web-react";
import Redis from "./blocks/redis";
import Postgres from "./blocks/postgre";
import CreateDatabase from "./create-database-editor";
import { useDatabase } from "@/api";


const { Row, Col } = Grid;
const { Title } = Typography;

const Overview = () => {

    const {database,isLoading} = useDatabase(); 

    return (
        <Card>
            <Typography.Title heading={5}>
                {'Database Overview'}
            </Typography.Title>

            <Divider />

            <Title heading={6}>Redis</Title>

            <Row gutter={20}>
                <Skeleton loading={isLoading} text={{ rows: 2, width: 60 }} animation>
                {
                    database
                        .filter((item) => item.databaseType === 'redis')
                        .map((item) => (
                                <Col key={`database-item-${item.id}`} span={6}>
                                    <Redis config={item} />
                                </Col> 
                            ))
                }
                </Skeleton>
            </Row>
            
            <Divider />

            <div>
                <Title heading={6}>Postgres</Title>

                <Row gutter={20}>
                    <Skeleton loading={isLoading} text={{ rows: 2, width: 60 }} animation>
                    {
                        database
                            .filter((item) => item.databaseType === 'postgres')
                            .map((item) => (
                                <>  
                                    <Col key={item.id} span={6}>
                                        <Postgres config={item} />
                                    </Col>
                                </>
                            ))
                    }
                    </Skeleton>
                </Row>
            </div>

                
            <Divider />

            <Title heading={6}>Database Management</Title>

            <CreateDatabase />
        </Card>
    )
}

export default Overview;