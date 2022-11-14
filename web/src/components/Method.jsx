import React from "react";

function Method(props) {
  const { method, setMethod } = props

  const onChange = (e) => {
    setMethod(e.target.value)
  }

  return (
    <select className="select select-bordered w-full max-w-xs" onChange={onChange} value={method}>
      <option value="GET">GET</option>
      <option value="POST">POST</option>
    </select>
  )
}

export default Method
