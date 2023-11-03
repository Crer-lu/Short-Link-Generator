import React,{ useState, useEffect } from "react";
import { Avatar, List, Col, Divider, Row} from 'antd';
import '../../style/infoindex.css'


export default function Infopage(props) {
    const [data, setData] = useState([]);
    const [pagenum, setPagenum] = useState(1);
    const [pagesize, setPagesize] = useState(3);
    const [maxflag,setMaxflag] = useState(0);

    const fetchData = () => {
        fetch(`http://localhost:8080/api/link/list?page_number=${pagenum}&page_size=${pagesize}`, {
            method: "GET",
            credentials: 'include',
            // headers: {
            //     'Content-Type': 'application/json',
            // },
        })
        .then(response => response.json())
        .then(data => {
            console.log(data)
            if (data.data.links)
            {
                setData(data.data.links);
                setMaxflag(0);
            }
                
            else{
                setMaxflag(1);
                // setPagenum(prevPageNum => prevPageNum - 1);
            }

        })
        .catch(error => console.error('Error fetching data:', error));
    };

    useEffect(() => {
        fetchData();
    }, [pagenum]);

    const handlePageUp = () => {
        if (pagenum > 1) {
            setPagenum(prevPageNum => prevPageNum - 1);
        }
    };

    const handlePageDown = () => {
        if(0===maxflag)
            setPagenum(prevPageNum => prevPageNum + 1);
    };

    return (
        <div className="big">
            <Row>
                <Col className="arrowbase" flex={2}>
                    <input className="arrow" type="button" value="↑" onClick={handlePageUp} />
                    <br />
                    <input className="arrow" type="button" value="↓" onClick={handlePageDown} />
                    <br />
                    <p className="pagetext" style={{display: maxflag ? 'none' : 'flex'}}>{pagenum ? pagenum : "/"}</p>
                    <p className="pagetext" style={{display: maxflag ? 'flex' : 'none'}}>已经是最后一页了</p>
                    <br />
                    <input type="button" className="quit-btn" value="quit" onClick={props.onClick}></input>

                    <br />
                    {/* <input type="button" className="quit-btn" value="test" onClick={()=>{
                                    //modify
                                    fetch(`http://localhost:8080/api/link/list?page_number=${127}&page_size=${5}`,{
                                        method:"GET",
                                        credentials: 'include'
                                    })
                                    .then(response=>response.json())
                                    .then(data=>{console.log(data);
                                        document.getElementById('deletemsg').innerText=data.msg//可能还要修改
                                    });
                                    //modified
                                    }}></input> */}
                </Col>
                <Col  flex={3}>
                    <List
                        itemLayout="horizontal"
                        dataSource={data}
                        renderItem={(item, index) => (
                            <List.Item>
                                <List.Item.Meta
                                    avatar={<Avatar src={`https://xsgames.co/randomusers/avatar.php?g=pixel&key=${index}`} />}
                                    title={<a href="https://ant.design">{item.title}</a>}
                                    description={
                                    <div>
                                        <a>short:{item.short}</a>
                                        <br />
                                        <a>origin:{item.origin}</a>
                                        <br />
                                        <a>comment:{item.comment}</a>
                                        <br />
                                        <a>start:{item.start_time}</a>
                                        <br />
                                        <a>end_time:{item.end_time}</a>
                                        <br />
                                        <a>active{item.active}</a>
                                    </div>
                                    }
                                />
                            </List.Item>
                        )}
                    />
                </Col>
            </Row>
        </div>
    );
}
