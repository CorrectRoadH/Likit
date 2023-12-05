import axios from "axios";
import React, { useState } from "react";
import { toast } from 'sonner'

const RedisEditor = () => {

    const [host, setHost] = useState("localhost");
    const [port, setPort] = useState("6379");
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
            toast.success("connect success");
        }).catch((err) => {
            toast.error(err.response.data.msg);
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