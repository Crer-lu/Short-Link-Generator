import React,{ useState, useEffect } from "react";
import { Avatar, List, Col, Divider, Row, Input, Form} from 'antd';
import '../../style/user.css'


export default function User(props) {
    const [id, setId] = useState(-1);
    const [name, setName] = useState("");
    const [email, setMail] = useState("");
    // const [pwd,setPwd] = useState(0);
    const [tomodify, setTomodify] = useState(0);

    const [getedemail, setGemail] = useState("");
    const [getedname, setGname] = useState("");


    const fetchInfo = () => {
        fetch(`http://localhost:8080/api/user/info`, {
            method: "GET",
            credentials: 'include',
        })
        .then(response => response.json())
        .then(data => {
            console.log(data);
            setId(data.data.id);
            setName(data.data.name)
            setMail(data.data.email)
        })
        .catch(error => console.error('Error fetching data:', error));
    };

    const fetch_modify = () => {
        console.log("getedemail")
        console.log(getedemail)
        console.log("getedname")
        console.log(getedname)
        if(getedemail){
            if(getedname){
                console.log(1)
                fetch(`http://localhost:8080/api/user/info`, {
                    method: "POST",
                    credentials: 'include',
                    headers:{
                        'Content-Type': 'application/json',
                    },
                    body:JSON.stringify({
                        "id": id,
                        "name": getedname,
                        "email": getedemail
                    })
                })
                .then(response => response.json())
                .then(data => {
                    console.log(data);
                });
            }
            else{
                console.log(2)
                fetch(`http://localhost:8080/api/user/info`, {
                    method: "POST",
                    credentials: 'include',
                    headers:{
                        'Content-Type': 'application/json',
                    },
                    body:JSON.stringify({
                        "id": id,
                        "name": name,
                        "email": getedemail
                    })
                })
                .then(response => response.json())
                .then(data => {
                    console.log(data);
                });
            }
        }
        else{
            if(getedname){
                console.log(3)
                fetch(`http://localhost:8080/api/user/info`, {
                    method: "POST",
                    credentials: 'include',
                    headers:{
                        'Content-Type': 'application/json',
                    },
                    body:JSON.stringify({
                        "id": id,
                        "name": getedname,
                        "email": email
                    })
                })
                .then(response => response.json())
                .then(data => {
                    console.log(data);
                });
            }
            else{
                console.log(4)
                fetch(`http://localhost:8080/api/user/info`, {
                    method: "POST",
                    credentials: 'include',
                    headers:{
                        'Content-Type': 'application/json',
                    },
                    body:JSON.stringify({
                        "id": id,
                        "name": name,
                        "email": email
                    })
                })
                .then(response => response.json())
                .then(data => {
                    console.log(data);
                });
            }
        }
        
    };

    const getEmail=(event)=>{
        setGemail(event.target.value)
    }

    const getName=(event)=>{
        setGname(event.target.value)
    }


    useEffect(() => {
        fetchInfo();
    }, []);
    
    //以下为修改密码模块

    const [oldPassword, setOldPassword] = useState("");
    const [newPassword, setNewPassword] = useState("");
    const [confirmPassword, setConfirmPassword] = useState("");
    const [disp, setDisp] = useState(0);

    const handleOldPasswordChange = (event) => {
        setOldPassword(event.target.value);
    };

    const handleNewPasswordChange = (event) => {
        setNewPassword(event.target.value);
    };

    const handleConfirmPasswordChange = (event) => {
        setConfirmPassword(event.target.value);
    };


    if(0===tomodify)
    return (
        <div className="big">
            <Row>
                <Col className="arrowbase" flex={2}>
                    <input className="arrow" type="button" value="user-info" />
                    <br />
                    <input className="arrow" type="button" value="modify-info" onClick={()=>
                        setTomodify(1)
                    }/>
                    <br />
                    <input className="arrow" type="button" value="modify-pwd" onClick={()=>
                        setTomodify(2)
                    }/>
                    <br />
                    <input type="button" className="quit-btn" value="quit" onClick={props.onClick}></input>
                </Col>
                <Col flex={4}>
                    <div className="info-container">
                        <div className="info">
                            <a className="label">Email</a>
                            <p className="value">{email}</p>
                        </div>
                        <div className="info">
                            <a className="label">Name</a>
                            <p className="value">{name}</p>
                        </div>
                        <div className="info">
                            <a className="label">ID</a>
                            <p className="value">{id}</p>
                        </div>
                    </div>
                </Col>
            </Row>
        </div>
    );
    else if (1===tomodify)
    return(
        <div className="big">
            <Row>
                <Col className="arrowbase" flex={2}>
                        <input className="arrow" type="button" value="user-info" onClick={()=>
                            {setTomodify(0);fetchInfo();}
                        }/>
                        <br />
                        <input className="arrow" type="button" value="modify-info" onClick={()=>
                            setTomodify(1)
                        }/>
                        <br />
                        <input className="arrow" type="button" value="modify-pwd" onClick={()=>
                            setTomodify(2)
                        }/>
                        <br />
                        <input type="button" className="quit-btn" value="quit" onClick={props.onClick}></input>
                </Col>
                <Col  flex={4}>
                    <div>
                        <a className="label">Email</a>
                        <input className="modify-input" onChange={getEmail}/>
                    </div>
                    <br />
                    <br />
                    <div>
                        <a className="label">Name</a>
                        <input className="modify-input" onChange={getName}/>
                    </div>
                    <br />
                    <input type="button" className="submit-btn" value="submit" onClick={fetch_modify}></input>
                    
                    
                </Col>
            </Row>
        </div>
        );
        else if (2===tomodify){
            return(
                <div className="big">
                    <Row>
                        <Col className="arrowbase" flex={2}>
                        <input className="arrow" type="button" value="user-info" onClick={()=>
                            {setTomodify(0);fetchInfo();}
                        }/>
                        <br />
                        <input className="arrow" type="button" value="modify-info" onClick={()=>
                            setTomodify(1)
                        }/>
                        <br />
                        <input className="arrow" type="button" value="modify-pwd" onClick={()=>
                            setTomodify(2)
                        }/>
                        <br />
                        <input type="button" className="quit-btn" value="quit" onClick={props.onClick}></input>
                        </Col>
                        <Col flex={4}>
                            <div className="change-password-container">
                                <div className="input-container">
                                    <a className="label">old password：</a>
                                    <br />
                                    <input
                                        type="password"
                                        style={{marginLeft:180}}
                                        value={oldPassword}
                                        onChange={handleOldPasswordChange}
                                    />
                                </div>
                                <div className="input-container">
                                    <a className="label">New password：</a>
                                    <br />
                                    <input
                                        type="password"
                                        style={{marginLeft:180}}
                                        value={newPassword}
                                        onChange={handleNewPasswordChange}
                                    />
                                </div>
                                <div className="input-container">
                                    <a className="label">confirm new password：</a>
                                    <br />
                                    <input
                                        type="password"
                                        style={{marginLeft:180}}
                                        value={confirmPassword}
                                        onChange={handleConfirmPasswordChange}
                                    />
                                </div>
                                <p style={{display: disp ? 'flex' : 'none',color:'red',fontSize:20}}>两次新密码输入不相同</p>
                                <input type="button" className="pwd-btn" value="submit" onClick={()=>{
                                    if (confirmPassword!=newPassword){
                                        setDisp(1);
                                    }
                                    else{
                                        setDisp(0);
                                        fetch("http://localhost:8080/api/user/passwd",{
                                            method: "POST",
                                            credentials: 'include',
                                            headers:{
                                                'Content-Type': 'application/json',
                                            },
                                            body:JSON.stringify({
                                                "new_pwd": confirmPassword,
                                                "old_pwd": oldPassword
                                            })
                                        })
                                        .then(response => response.json())
                                        .then(data => {
                                            console.log(data);
                                        })
                                    }
                                }}/>
                            </div>
                            
                        </Col>
                    </Row>
                </div>
                );
        }
}
