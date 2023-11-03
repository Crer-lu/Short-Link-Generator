import React from "react";
import Register from "./register";
import Shortlink from "./shortlink";
import "../style/loginandregister.css"
import myImage from '../assets/ironman.jpg';

class Login extends React.Component{
    constructor(props){
        super(props);
        this.state={
            getedname:"",
            getedpwd:"",
            getedmail:"",
            responcode:0,
            responjwt:"",
            catchaUrl: "",
            catchaid:"",
            getedCap:"",
            respon:0,
            message:""
        };
    }

    componentDidMount() {
        // 在组件渲染后请求 catcha 的初始 URL
        this.fetchCatchaUrl();
      }
    
    fetchCatchaUrl() {
        fetch("http://localhost:8080/api/user/captcha",{//modify
            method:"GET",
            credentials: 'include',
        }) 
        .then((response) => response.json())
        .then((data) => {
            console.log(data.data.captcha_url);
            console.log("aaa");
            this.setState({
            catchaUrl: "http://localhost:8080"+data.data.captcha_url,
            catchaid:data.data.captcha_id
            });
        });
    }
    

    getName=(event)=>{
        this.setState({
            getedname:event.target.value
        })
    }

    getMail=(event)=>{
        this.setState({
            getedmail:event.target.value
        })
    }

    getPwd=(event)=>{
        this.setState({
            getedpwd:event.target.value
        })
    }

    getCaptcha=(event)=>{
        this.setState({
            getedCap:event.target.value
        })
    }

    handleClick(){
        console.log(this.state.responcode)
        console.log(this.state.responjwt)
        fetch("http://localhost:8080/api/user/logout",{//modify
            method:"POST",
            credentials: 'include',
        })
        .then(response=>response.json())
        .then(data=>console.log(data))
        this.setState({
            responcode:0
        })
    }

    render(){
        if(0===this.state.responcode)
        return(
            <div className="container">
                <div className="box">
                    <div className="title">Register</div>
                    <Register/>
                </div>
                
                <div className="box">
                    {/* <img src={myImage} alt="Your Image" class="catcha"/> */}
                    <img
                        src={this.state.catchaUrl} // 使用 catchaUrl
                        alt="Your Image"
                        className="catcha"
                        onClick={() => this.fetchCatchaUrl()} // 在点击时请求新的 catcha URL
                    />
                    <div className="title">Login</div>
                    {/* <input type="text" name="name" id="logname" className="input-text" placeholder="username" onChange={this.getName}></input><br /> */}
                    <input type="text" name="name" id="logmail" className="input-text" placeholder="email" onChange={this.getMail}></input><br />
                    <input type="password" name="password" id="logpwd" className="input-text" placeholder="password" onChange={this.getPwd}></input><br />
                    <input type="text" name="captcha" className="input-text" placeholder="captcha" onChange={this.getCaptcha}></input><br />
                    <input type="button" className="input-button" value="LOGIN" onClick={()=>{
                        this.fetchCatchaUrl();
                        fetch("http://localhost:8080/api/user/login",{
                            method:"POST",
                            credentials: 'include',
                            headers:{'Content-Type':'application/json'},
                            body:JSON.stringify({
                                // "name": this.state.getedname,
                                "email":this.state.getedmail,
                                "password": this.state.getedpwd,
                                "captcha_id":this.state.catchaid,
                                "captcha_value":this.state.getedCap
                            })
                        })
                        .then(response=>response.json())
                        .then(data=>{console.log(data);
                                console.log(data.code);
                                this.setState({
                                    responcode:data.code,
                                    message:data.msg,
                                    respon:1,
                                });
                                // console.log(this.state.responcode)这里会显示0，要整个onClick结束以后responcode才会变
                            })
                        // document.getElementById("logname").value="";
                        // document.getElementById("logcap").value="";
                        // document.getElementById("logmail").value="";
                        // document.getElementById("logpwd").value="";
                    }}>
                    </input>
                </div>
            </div>
        );
        else if(200===this.state.responcode){
            return (
                <Shortlink onClick={()=>this.handleClick()} responjwt={this.state.responjwt}/>
            );
        }
        else{
            return(
                <div className="container">
                <div className="box">
                    <div className="title">Register</div>
                    <Register/>
                </div>
                <div className="box">
                    {/* <div className="title">Login</div>
                    <div>Email or password is wrong.</div> */}
                    {/* <input type="text" name="name" id="logname" className="input-text" placeholder="username" onChange={this.getName}></input><br /> */}
                    {/* <input type="text" name="name" id="logmail" className="input-text" placeholder="email" onChange={this.getMail}></input><br />
                    <input type="text" className="input-text_"></input><br />
                    <input type="password" name="password" id="logpwd" className="input-text" placeholder="password" onChange={this.getPwd}></input><br />
                    <input type="button" className="input-button" value="LOGIN" onClick={()=>{
                        console.log(this.state.getedname);
                        console.log(this.state.getedpwd);
                        fetch("http://localhost:3000/user/login",{
                            method:"POST",
                            headers:{'Content-Type':'application/json'},
                            body:JSON.stringify({
                                // "name": this.state.getedname,
                                "email":this.state.getedmail,
                                "password": this.state.getedpwd
                            })
                        })
                        .then(response=>response.json())
                        .then(data=>{console.log(data);
                                console.log(data.code);
                                this.setState({
                                    responcode:data.code,
                                    responjwt:data.jwt
                                });
                                // console.log(this.state.responcode)这里会显示0，要整个onClick结束以后responcode才会变
                            })
                        // document.getElementById("logname").value="";
                        document.getElementById("logmail").value="";
                        document.getElementById("logpwd").value="";
                    }}>
                    </input> */}
                    <img
                        src={this.state.catchaUrl} // 使用 catchaUrl
                        alt="Your Image"
                        className="catcha"
                        onClick={() => this.fetchCatchaUrl()} // 在点击时请求新的 catcha URL
                    />
                    <div className="title">Login</div>
                    {/* <input type="text" name="name" id="logname" className="input-text" placeholder="username" onChange={this.getName}></input><br /> */}
                    <input type="text" name="name" id="logmail" className="input-text" placeholder="email" onChange={this.getMail}></input><br />
                    <input type="password" name="password" id="logpwd" className="input-text" placeholder="password" onChange={this.getPwd}></input><br />
                    <input type="text" name="captcha" id="logcap" className="input-text" placeholder="captcha" onChange={this.getCaptcha}></input><br />
                    <input type="button" className="input-button" value="LOGIN" onClick={()=>{
                        this.fetchCatchaUrl();
                        // this.setState({responcode:200})
                        fetch("http://localhost:8080/api/user/login",{
                            method:"POST",
                            credentials: 'include',
                            headers:{'Content-Type':'application/json'},
                            body:JSON.stringify({
                                // "name": this.state.getedname,
                                "email":this.state.getedmail,
                                "password": this.state.getedpwd,
                                "captcha_id":this.state.catchaid,
                                "captcha_value":this.state.getedCap
                            })
                        })
                        .then(response=>response.json())
                        .then(data=>{console.log(data);
                                console.log(data.code);
                                this.setState({
                                    responcode:data.code,
                                    message:data.msg,
                                    respon:1,
                                });
                                // console.log(this.state.responcode)这里会显示0，要整个onClick结束以后responcode才会变
                            })
                        // document.getElementById("logcap").value="";
                        // document.getElementById("logmail").value="";
                        // document.getElementById("logpwd").value="";
                    }}>
                    </input>
                    <p style={{display: this.state.respon ? 'flex' : 'none',color:'red'}}>{this.state.message}</p>
                </div>
            </div>
            );
        }
    }
}

export default Login