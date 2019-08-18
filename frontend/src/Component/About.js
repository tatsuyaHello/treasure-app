import React, { Component } from "react";
import { BrowserRouter, Route, Link } from 'react-router-dom'

import { showLecture } from "../API/api"

// export const About = ({ match }) => (
//   <div>
//     <h2>About</h2>
//     <p>フレンズに投票するページです</p>
//     <p>DEMO: {match.params.id}</p>
//   </div>
// )

// export class About extends Component {
//   render() {
//     const {params} = this.props.match
//     return(
//       <div>
//         <h2>About</h2>
//         <p>フレンズに投票するページです</p>
//         <p>DEMO: {params.id}</p>
//       </div>
//     )
//   }
// }


export class About extends Component {
  constructor(props) {
    super(props);
    this.state = {
      result_title: '',
      result_english_title: '',
      result_unit: '',
      result_semester: '',
      result_location: '',
      result_lecture_style: '',
      result_teacher: '',
      result_overview: '',
      result_goal: '',
      result_sub_title: '',
      result_lecture_id: '',
      result_textbook: '',
      result_remarks: '',
      scehdule_list: [],
      evaluate_list: []
    };
  }

  componentDidMount() {
    // event.preventDefault();
    let res = this.getLecture();
    console.log(res)
    res.then(data => {
      this.setState({
        result_title:data.title,
        result_english_title:data.english_title,
        result_unit:data.unit,
        result_semester:data.semester,
        result_location:data.location,
        result_lecture_style:data.lecture_style,
        result_teacher:data.teacher,
        result_overview:data.overview,
        result_goal:data.goal,
        result_sub_title:data.sub_title,
        result_lecture_id:data.lecture_id,
        result_textbook:data.textbook,
        result_remarks:data.remarks,
        scehdule_list:data.Scehdule,
        evaluate_list: data.Evaluate
      })
    })
  }
  
  getLecture() {
    const {params} = this.props.match
    console.log(params)
    // event.preventDefault();
    console.log(params.id)
    let res = showLecture(params.id);
    return res
  }

  render() {
    var all = {
      marginTop: 30,
      marginBottom: 30,
    };
    var head = {
      textAlign: "left",
      fontSize: 20,
      marginRight: 30,
      marginLeft: 30,
      margin: "auto",
      width: 800,
      border: "ridge",
      borderWidth: "thin",
      borderColor: "#008080",
      borderRadius: 5,
    };
    var header = {
      height: 60,
      backgroundColor: "#008080",
      fontSize: 24,
      color: "white",
      textAlign: "center"
    };
    var mainTitle = {
      width: 200,
      marginTop: 10,
      float: "left"
    };
    var reviewButton = {
      width: 140,
      marginTop: 16,
      marginRight: 4,
      fontSize: 12,
      float: "right",
      backgroundColor: "#FF9872",
      borderColor: "#008080",
      borderRadius: 5,
      padding: 6
    };
    var headTable = {
      textAlign: "center"
    };
    var tableHead = {
      fontSize: 16,
      paddingRight: 20,
      paddingLeft: 20,
      textAlign: "left"
    };
    var tableData = {
      fontSize: 16,
      paddingRight: 20,
      paddingLeft: 20,
      paddingBottom: 10,
      textAlign: "left"
    };
    var mainBody = {
      width: 800,
      margin: "auto",
    };
    var middle = {
      display: "-webkit-flex",
      display: "flex",
      webkitFlexDirection: "row",
      flexDirection: "row",
      webkitFlexWrap: "wrap",
      flexWrap: "wrap",
    };
    var middleLeft = {
      width: 360,
      border: "ridge",
      borderWidth: "thin",
      borderRadius: 5,
      paddingLeft: 10,
      marginTop: 30,
      fontSize: 14
    };
    var column = {
      color: "#007700",
      fontWeight: "bold",
    }
    var middleRight = {
      width: 360,
      border: "ridge",
      borderWidth: "thin",
      borderRadius: 5,
      paddingLeft: 10,
      marginTop: 30,
      marginLeft: 50,
      fontSize: 14
    };
    var buttom = {
      display: "-webkit-flex",
      display: "flex",
      webkitFlexDirection: "row",
      flexDirection: "row",
      webkitFlexWrap: "wrap",
      flexWrap: "wrap",
    };
    var buttomLeft = {
      width: 350,
      border: "ridge",
      borderWidth: "thin",
      borderRadius: 5,
      paddingLeft: 10,
      paddingRight: 10,
      paddingTop: 10,
      marginTop: 30,
      fontSize: 14
    };
    var lectureDetail = {
      fontSize: 12,
    }
    var buttomRight = {
      width: 350,
      border: "ridge",
      borderWidth: "thin",
      borderRadius: 5,
      paddingLeft: 10,
      paddingRight: 10,
      paddingTop: 10,
      marginTop: 30,
      marginLeft: 50,
      fontSize: 14
    };
    var scehduleDetail = {
      
    }
    return (
      <div style={all}>
        <div style={head}>
          <div style={header}>
            <div style={mainTitle}>
             {this.state.result_title}
            </div>
            <div style={reviewButton}>
             レビューを書く
            </div>
          </div>
          <div style={headTable}>
            <tbody>
              <tr>
                <th style={tableHead}>年度</th>
                <th style={tableHead}>講義番号</th>
                <th style={tableHead}>講師</th>
                <th style={tableHead}>学期</th>
                <th style={tableHead}>単位</th>
                <th style={tableHead}>キャンパス</th>
                <th style={tableHead}>講義スタイル</th>
              </tr>
              <tr>
                <td style={tableData}>2019</td>
                <td style={tableData}>{this.state.result_lecture_id}</td>
                <td style={tableData}>{this.state.result_teacher}</td>
                <td style={tableData}>{this.state.result_semester}</td>
                <td style={tableData}>{this.state.result_unit}</td>
                <td style={tableData}>{this.state.result_location}</td>
                <td style={tableData}>{this.state.result_lecture_style}</td>
              </tr>
            </tbody>
          </div>
        </div>
        <div style={mainBody}>
          <div style={middle}>
            <div style={middleLeft}>
              <div>
                <p style={column}>科目コード/科目名</p>
              </div>
              <div>
                <p>{this.state.result_lecture_id}/{this.state.result_title}</p>
              </div>
              <div>
                <p style={column}>Englishタイトル</p>
              </div>
              <div>
                <p>{this.state.result_english_title}</p>
              </div>
              <div>
                <p style={column}>担当者</p>
              </div>
              <div>
                <p>{this.state.result_teacher}</p>
              </div>
            </div>
            <div style={middleRight}>
              <div>
                <p></p>
              </div>
              <div>
                {this.state.evaluate_list.map((l)=>(
                  <div>
                    {l.method}
                    {l.percentage}
                    {l.comment}
                  </div>
                  ))}
              </div>
            </div>
          </div>
          <div style={buttom}>
            <div style={buttomLeft}>
              <div style={column}>
                授業ゴール
              </div>
              <div style={lectureDetail}>
                {this.state.result_goal}
              </div>
              <div style={column}>
                授業内容
              </div>
              <div style={lectureDetail}>
                {this.state.result_overview}
              </div>
            </div>
            <div style={buttomRight}>
              <div style={column}>
                授業計画
              </div>
              <div style={scehduleDetail}>
                {this.state.scehdule_list.map((l, num)=>(
                  <div>
                    {num.number}: {l.session}
                  </div>
                  ))}
              </div>
            </div>
          </div>
        </div>
      </div>
    )
  }
}
