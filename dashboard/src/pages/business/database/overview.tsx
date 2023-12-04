import React, { useState } from "react";
import RedisEditor from "./redis-editor";
import { Card, Divider, Typography } from "@arco-design/web-react";



const Overview = () => {
    const [data, setData] = useState<any[]>([]);
    
    return (
        <Card>
            <Typography.Title heading={5}>
                {'Database Overview'}
            </Typography.Title>

            <div>
                redis
            </div>

            <div>
                postgres
            </div>

                
            <Divider />

            <div>new source</div>

            <RedisEditor />
        </Card>
    )
}

export default Overview;