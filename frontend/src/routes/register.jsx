import React from "react";
import { json } from "react-router-dom";

class Register extends React.Component{
    state={
        getedname:"",
        getedemail:"",
        getedpwd:"",
        shouldHide:1,
        areyouOK:0,
        msg:""
    };



    getName=(event)=>{
        this.setState({
            getedname:event.target.value
        })
    }

    getEmail=(event)=>{
        this.setState({
            getedemail:event.target.value
        })
    }

    getPwd=(event)=>{
        this.setState({
            getedpwd:event.target.value
        })
    }

    render(){
        return(
            <div>
                <input type="text" name="name" id="rgname" className="input-text" placeholder="username" onChange={this.getName}></input><br />
                <input type="text" name="email" id="rgemail" className="input-text" placeholder="email" onChange={this.getEmail}></input><br />
                <input type="text" name="password" id="rgpwd" className="input-text" placeholder="password" onChange={this.getPwd}></input><br />
                
                <input
                    type="button"
                    className="input-button"
                    value="REGISTER"
                    onClick={() => {
                        const email = this.state.getedemail;
                        this.setState({shouldHide:1})
                        if (
                            email.includes("@") && // 检查是否包含 @ 符号
                            email.indexOf("@") !== 0 && // 检查 @ 是否不是第一个字符
                            email.indexOf("@") !== email.length - 1 // 检查 @ 是否不是最后一个字符
                        ) {
                            console.log(this.state.getedname);
                            console.log(this.state.getedemail);
                            console.log(this.state.getedpwd);

                            fetch("http://localhost:8080/api/user/register", {
                                method: "POST",
                                credentials: 'include',
                                mode: 'cors', // 设置为 'cors'
                                headers: { 'Content-Type': 'application/json' },
                                body: JSON.stringify({
                                    "name": this.state.getedname,
                                    "email": this.state.getedemail,
                                    "password": this.state.getedpwd
                                })
                            })
                            .then(response => {
                                console.log(response);
                                return response.json();
                            })
                            .then(data => {
                                console.log(data.msg);
                                this.setState({
                                    areyouOK:1,
                                    msg:data.msg
                                });
                            });
                        } else {
                            this.setState({
                                shouldHide:0
                            })
                        }

                        document.getElementById("rgname").value = "";
                        document.getElementById("rgemail").value = "";
                        document.getElementById("rgpwd").value = "";
                    }}
                />
                <p style={{display: this.state.shouldHide ? 'none' : 'flex',color:'red'}}>邮箱格式错误</p>
                <p style={{display: this.state.areyouOK ? 'flex' : 'none',color:'red'}}>{this.state.msg}</p>
            </div>
        );
    }
}

export default Register