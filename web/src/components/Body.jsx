import React, { useState } from "react";

function Body(props) {
  const { body, setBody } = props
  const [value, setValue] = useState(body)

  const onClick = () => {
    setBody(value)
  }

  return (
    <div className="flex h-auto mt-5">
      <input type="text" onChange={(e) => setValue(e.target.value)} value={value}
             placeholder="Request Body" className="grow input input-bordered"/>
      <button className="flex-none btn btn-accent ml-2.5" onClick={onClick}>Apply</button>
    </div>
  )
}

export default Body
