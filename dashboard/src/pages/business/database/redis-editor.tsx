import axios from "axios";
import React, { useState } from "react";

const RedisEditor = () => {

    const [host, setHost] = useState("");
    const [port, setPort] = useState("");
    const [password, setPassword] = useState("");

    const handleTestBtnClick = () => {
        axios.post("/admin/v1/database/test", {
            "databaseType":"redis",
            "host": host,
            "port": Number(port),
            "username":"",
            "password": password,
            "database":""
        }).then((res) => {
            console.log(res);
        })
    }

    return (
        <div>
            <div>host:</div>
            <input value={host} onChange={(e)=>setHost(e.target.value)} />

            <div>port:</div>
            <input value={port} onChange={(e)=>setPort(e.target.value)} />

            <div>password:</div>
            <input value={password} onChange={(e)=>setPassword(e.target.value)} />
        
            <button
                onClick={handleTestBtnClick}
            >test connect</button>
        </div>
    )
}

export default RedisEditor;