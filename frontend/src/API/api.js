import React from "react";

export const searchLecture = async function(state)  {
  if (state.title === "" && state.english_title === "" && state.unit === "" && state.semester === "" && state.location === "" && state.lecture_style === "" && state.teacher === "" && state.overview === "") {
    return fetch("http://localhost:1991/lectures", {
      method: "GET",
      headers: new Headers({
        "Content-Type": "application/json",
      }),
    }).then(res => {
      if (res.ok) {
        return res.json();
      } else {
        throw Error(`Request rejected with status ${res.status}`);
      }
    })
  } else {
    let url = "http://localhost:1991/lecture?"+"title="+state.title+"&english_title="+state.english_title+"&unit="+state.unit+"&semester="+state.semester+"&location="+state.location+"&lecture_style="+state.lecture_style+"&teacher="+state.teacher+"&overview="+state.overview
    console.log(url)
    return fetch(url, {
      method: "GET",
      headers: new Headers({
        "Content-Type": "application/json",
      }),
    }).then(res => {
      if (res.ok) {
        return res.json();
      } else {
        throw Error(`Request rejected with status ${res.status}`);
      }
    })
  }
}
