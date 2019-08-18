import React from "react";
import { BrowserRouter, Route, Link } from 'react-router-dom'

import { searchLecture } from "../API/api"


export class Form extends React.Component {

  constructor(props) {
    super(props);
    this.state = {
      title: '',
      english_title: '',
      unit: '',
      semester: '',
      location: '',
      lecture_style: '',
      teacher: '',
      overview: '',
      list: [],
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

    this.handleTitleChange = this.handleTitleChange.bind(this);
    this.handleEnglishTitleChange = this.handleEnglishTitleChange.bind(this);
    this.handleUnitChange = this.handleUnitChange.bind(this);
    this.handleSemester = this.handleSemester.bind(this);
    this.handleLocation = this.handleLocation.bind(this);
    this.handleLectureStyle = this.handleLectureStyle.bind(this);
    this.handleTeacher = this.handleTeacher.bind(this);
    this.handleOverview = this.handleOverview.bind(this);
    this.handleSubmit = this.handleSubmit.bind(this)
  }


  handleTitleChange(event) {
    this.setState({title: event.target.value});
  }
  handleEnglishTitleChange(event) {
    this.setState({english_title: event.target.value});
  }
  handleUnitChange(event) {
    this.setState({unit: event.target.value});
  }
  handleSemester(event) {
    this.setState({semester: event.target.value});
  }
  handleLocation(event) {
    this.setState({location: event.target.value});
  }
  handleLectureStyle(event) {
    this.setState({lecture_style: event.target.value});
  }
  handleTeacher(event) {
    this.setState({teacher: event.target.value});
  }
  handleOverview(event) {
    this.setState({overview: event.target.value});
  }

  handleSubmit(event) {
    event.preventDefault();
    let res = searchLecture(this.state)
    console.log(res)
    res.then(data => {
      this.setState({list: []});
      for (let i = 0;i < data.length; i++) {
        const list = [];
        for(let content of data) {
          list.push(
            {
              result_title: content.title,
              result_english_title: content.english_title,
              result_unit: content.unit,
              result_semester: content.semester,
              result_location: content.location,
              result_lecture_style: content.lecture_style,
              result_teacher: content.teacher,
              result_overview: content.overview,
              result_goal: content.goal,
              result_sub_title: content.sub_title,
              result_lecture_id: content.lecture_id,
              result_textbook: content.textbook,
              result_remarks: content.remarks,
              scehdule_list: content.Scehdule,
              evaluate_list: content.Evaluate
            }
          )
        }
        this.setState({list: list});
      }
    })
  }

  render() {
    var search = {
      textAlign: "left",
      fontSize: 20,
      marginRight: 30,
      marginLeft: 30,
      margin: "auto",
      width: 800,
      border: "ridge",
      borderWidth: "thin",
      borderColor: "#337AB7",
      borderRadius: 5
    };
    var searchHeader = {
      height: 40,
      backgroundColor: "#337AB7",
      color: "white",
      fontSize: 16,
      textAlign: "center"
    };
    var searhAllForm = {
      paddingBottom: 30
    };
    var searchForm = {
      margin: 5,
      fontSize: 30
    };
    var detailSearchForm = {
      padding: 10,
      border: "ridge",
      borderWidth: "thin",
      borderRadius: 5
    };
    var submitForm = {
      marginLeft: 200,
      marginTop: 30,
      paddingRight: 10,
      paddingLeft: 10,
      fontSize: 16,
      color: "white",
      border: "ridge",
      borderWidth: "thin",
      borderRadius: 5,
      backgroundColor: "#E36717"
    };
    var searchResult={
      fontSize: 12,
      color: "#444444",
      marginRight: 30,
      marginLeft: 30,
      margin: "auto",
      width: 800,
      display: "-webkit-flex",
      display: "flex",
      webkitFlexDirection: "row",
      flexDirection: "row",
      webkitFlexWrap: "wrap",
      flexWrap: "wrap",
    };
    var searchResultUnit={
        width: 230,
        border: "ridge",
        borderWidth: "thin",
        borderRadius: 5,
        margin: 10,
        padding: 3,
    };
    var goal={
      fontSize: 8
    };
    return (
      <div>
        <div style={search}>
          <div style={searchHeader}>
            検索項目
          </div>
          <form onSubmit={this.handleSubmit} style={searhAllForm}>
            <label style={searchForm}>
              <input type="text" value={this.state.title} placeholder="講義名" style={detailSearchForm} onChange={this.handleTitleChange} />
            </label>
            <label style={searchForm}>
              <input type="text" value={this.state.english_title} placeholder="英語名" style={detailSearchForm} onChange={this.handleEnglishTitleChange} />
            </label>
            <label style={searchForm}>
              <input type="text" value={this.state.unit} placeholder="単位数" style={detailSearchForm} onChange={this.handleUnitChange} />
            </label>
            <label style={searchForm}>
              <input type="text" value={this.state.semester} placeholder="学期" style={detailSearchForm} onChange={this.handleSemester} />
            </label>
            <label style={searchForm}>
              <input type="text" value={this.state.location} placeholder="キャンパス" style={detailSearchForm} onChange={this.handleLocation} />
            </label>
            <label style={searchForm}>
              <input type="text" value={this.state.lecture_style} placeholder="講義スタイル" style={detailSearchForm} onChange={this.handleLectureStyle} />
            </label >
            <label style={searchForm}>
              <input type="text" value={this.state.teacher} placeholder="講師" style={detailSearchForm} onChange={this.handleTeacher} />
            </label>
            <label style={searchForm}>
              <input type="text" value={this.state.overview} placeholder="概要" style={detailSearchForm} onChange={this.handleOverview} />
            </label>
            <label ></label>
            <input type="submit" style={submitForm} value="検索" />
          </form>
        </div>
        <div style={searchResult}>
            {this.state.list.map((l)=>(
                    <div key={l.result_title} style={searchResultUnit}>
                      <div>
                        <Link style={{ textDecoration: 'none' }} to='/about'>{l.result_lecture_id}</Link> / {l.result_unit}
                      </div>
                      <div>
                        {l.result_title} {l.result_location}
                      </div>
                      <div>
                        {l.result_semester} {l.result_teacher}
                      </div>
                      <br />
                      <div style={goal}>
                        {l.result_goal}
                      </div>
                      <div>----------------------------------------</div>
                      <br />
                      <div>
                        {l.evaluate_list.map((evalu)=>(
                          <div>{evalu.method} {evalu.comment} {evalu.percentage}</div>
                          ))}
                      </div>
                    </div>
              ))
            }
        </div>
      </div>
    );
  }
}
