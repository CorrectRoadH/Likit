import React from "react";
import RedisEditor from "./redis-editor";

const Overview = () => {
    return (
        <div>
            <div>data source</div>
            <div>
                redis
            </div>
            
            <div>new source</div>

            <RedisEditor />
        </div>
    )
}

export default Overview;