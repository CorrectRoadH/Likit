import React, { useEffect, useState } from "react";
import RedisEditor from "./editor/redis-editor";
import { Card, Divider, Typography,Grid } from "@arco-design/web-react";
import { DatabaseConnectionConfig } from "./types";
import axios from "axios";
import Redis from "./blocks/redis";
import Postgres from "./blocks/postgre";
import CreateDatabase from "./create-database-editor";


const { Row, Col } = Grid;

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

            <div>
                redis
            </div>
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
                postgres

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

            <div>new source</div>

            <CreateDatabase />
        </Card>
    )
}

export default Overview;