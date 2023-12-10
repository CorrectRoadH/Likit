import React, { useEffect, useState } from "react";
import { Card, Divider, Typography,Grid } from "@arco-design/web-react";
import { DatabaseConnectionConfig } from "../../../types";
import axios from "axios";
import Redis from "./blocks/redis";
import Postgres from "./blocks/postgre";
import CreateDatabase from "./create-database-editor";


const { Row, Col } = Grid;
const { Title } = Typography;

const Overview = () => {

    // TODO extract to a hook
    const [data, setData] = useState<DatabaseConnectionConfig[]>([]);
    
    const fetchData = () => {
        axios.get('/admin/v1/database').then((res) => {
            setData(res.data.dataSourceConfig);
        });
    }

    useEffect(() => {
        fetchData()
    }, []);
    return (
        <Card>
            <Typography.Title heading={5}>
                {'Database Overview'}
            </Typography.Title>

            <Divider />

            <Title heading={6}>Redis</Title>

            <Row>
                {
                    data
                        .filter((item) => item.databaseType === 'redis')
                        .map((item) => (
                                <>  
                                <Col key={item.id} span={6}>
                                    <Redis config={item} />
                                </Col>
                                </>
 
                            ))
                }
            </Row>
            
            <Divider />

            <div>
                <Title heading={6}>Postgres</Title>

                <Row>
                {
                    data
                        .filter((item) => item.databaseType === 'postgres')
                        .map((item) => (
                            <>  
                                <Col key={item.id} span={6}>
                                    <Postgres config={item} />
                                </Col>
                            </>
                        ))
                }
            </Row>
            </div>

                
            <Divider />

            <Title heading={6}>Database Management</Title>

            <CreateDatabase />
        </Card>
    )
}

export default Overview;