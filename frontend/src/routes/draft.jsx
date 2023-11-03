import React from "react";
import '../style/shortlink.css'

class Shortlink extends React.Component{

    constructor(props){
        super(props);
        this.state={
            getedorigin:"",
            getedshort:"",
            myform:"",
            linkid:"",
            shorted:""
        }
    }

    getOrigin=(event)=>{
        this.setState({
            getedorigin:event.target.value
        })
    }

    getShort=(event)=>{
        this.setState({
            getedshort:event.target.value
        })
    }

    render(){
        return(
            <div className="mainbody">
                <div className="text">
                    <div style={{paddingLeft:100,paddingBottom:15}}><h2 style={{marginBottom:5}}>Shorten and personalize any link.</h2></div>
                    <div style={{paddingLeft:100}}><h3>Get real-time traffic statistics for your links.</h3></div>
                    <div style={{paddingLeft:100,paddingBottom:15}}><h3>Free service.</h3></div>
                    <div className="contain">
                    {/* <ol> */}
                        {/* <li> */}
                        <div className="case">
                            <div className="caseinside">
                            <p>Link to shorten</p>
                            <input type="text" className="input-txt" id="linktoshort" onChange={this.getOrigin}></input>
                            <p>Customize your short link(optional):</p>
                            localhost:3000/bit.do/<input type="text" id="diy" className="input-txt_" onChange={this.getShort}></input>
                            <br />
                            <input type="button" className="input-btn" value="Shorten" onClick={()=>{
                                    this.state.myform = new FormData();
                                    this.state.myform.append("origin",this.state.getedorigin);
                                    this.state.myform.append("short",this.state.getedshort);
                                    fetch("http://localhost:3000/url/create",{
                                        method:"POST",
                                        headers:{"Authorization":"Bearer "+ this.props.responjwt},
                                        body:this.state.myform
                                    })
                                    .then(response=>response.json())
                                    .then(data=>{console.log(data);
                                        this.setState({
                                            linkid:data.urlinfo.id,
                                            shorted:data.urlinfo.short
                                        });
                                        document.getElementById('generated').value='localhost:3000/'+data.urlinfo.short;
                                    });
                                    document.getElementById("linktoshort").value="";
                                    document.getElementById("diy").value="";
                                    }
                                }>
                            </input>
                            </div>
                        </div>
                        {/* </li>
                        <li> */}
                        <div className="case">
                        <div>
                            <p>Generated Link</p><br />
                            <p>ID:{this.state.linkid}</p><br/>
                            <p>Generated Link:(valid for 24h)</p>
                            <input type="text" className="input-txt" id="generated"></input>
                            <br />
                        </div>
                        </div>
                        {/* </li>
                        <li> */}
                        <div className="case">
                        <div>
                            input id to pause certain short link:<br/>
                            <input type="text" className="input-txt" id="pause"></input><br/>
                            <input type="button" className="input-btn" value="pause" onClick={()=>{
                                    this.state.myform = new FormData();
                                    this.state.myform.append("id",document.getElementById('pause').value);
                                    fetch("http://localhost:3000/url/pause",{
                                        method:"POST",
                                        headers:{"Authorization":"Bearer "+ this.props.responjwt},
                                        body:this.state.myform
                                    })
                                    .then(response=>response.json())
                                    .then(data=>{console.log(data);
                                        document.getElementById('pausemsg').innerText=data.msg
                                    });
                                    }
                                }>
                            </input>
                            <p id="pausemsg" style={{fontSize:10, color:'red'}}></p>
                        </div>
                        </div>
                        {/* </li>
                        <li> */}
                        <div className="case">
                        <div>
                            input id to delete certain short link:<br/>
                            <input type="text" className="input-txt" id="delete"></input><br/>
                            <input type="button" className="input-btn" value="delete" onClick={()=>{
                                    this.state.myform = new FormData();
                                    this.state.myform.append("id",document.getElementById('delete').value);
                                    fetch("http://localhost:3000/url/delete",{
                                        method:"DELETE",
                                        headers:{"Authorization":"Bearer "+ this.props.responjwt},
                                        body:this.state.myform
                                    })
                                    .then(response=>response.json())
                                    .then(data=>{console.log(data);
                                        document.getElementById('deletemsg').innerText=data.msg
                                    });
                                    }
                                }>
                            </input>
                            <p id="deletemsg" style={{fontSize:10, color:'red'}}></p>
                        </div>
                        </div>
                        {/* </li>
                        <li> */}
                        <div className="case">
                        <div>
                            <p>update certain shortlink</p>
                            <input type="text" className="input-txt" id="update" placeholder="please input the link ID"></input><br/>
                            <input type="text" className="input-txt" id="newshort" placeholder="please input the new shorted link"></input><br/>
                            <input type="text" className="input-txt" id="comment" placeholder="whatever you want to say(optional)"></input><br/>
                            <input type="button" className="input-btn" value="update"onClick={()=>{
                                    this.state.myform = new FormData();
                                    this.state.myform.append("id",document.getElementById('update').value);
                                    this.state.myform.append("newshort",document.getElementById('newshort').value);
                                    this.state.myform.append("comment",document.getElementById('comment').value);
                                    fetch("http://localhost:3000/url/update",{
                                        method:"PUT",
                                        headers:{"Authorization":"Bearer "+ this.props.responjwt},
                                        body:this.state.myform
                                    })
                                    .then(response=>response.json())
                                    .then(data=>{console.log(data)
                                    });
                                    }
                                }>
                            </input>
                        </div>
                        </div>
                        {/* </li>
                        <li> */}
                        <div className="case">
                        <div>
                            <input type="text" className="input-txt" placeholder="please input the link ID" id="query"></input>
                            <input type="button" className="input-btn" value="query"onClick={()=>{
                                    this.state.myform = new FormData();
                                    this.state.myform.append("id",21);
                                    
                                    fetch("http://localhost:3000/url/query",{
                                        method:"POST",
                                        headers:{"Authorization":"Bearer "+ this.props.responjwt},
                                        body:this.state.myform
                                    })
                                    .then(response=>response.json())
                                    .then(data=>{console.log(data)
                                    });
                                    }
                                }>
                            </input>
                            <p id="querymsg" style={{fontSize:10, color:'blue'}}></p>
                        {/* </li> */}
                        </div>
                        </div>
                    {/* </ol> */}
                    </div>
                    <input type="button" className="input-btn" value="logout" onClick={()=>this.props.onClick()}></input>
                </div>
            </div>
        );
    }
}

export default Shortlink