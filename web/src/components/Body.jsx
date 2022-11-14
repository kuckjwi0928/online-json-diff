import React, { useState } from "react";

function Body(props) {
  const { addBody } = props
  const [bodyKey, setBodyKey] = useState('');
  const [bodyValue, setBodyValue] = useState('');

  const onClick = () => {
    if (!bodyKey || !bodyValue) {
      return
    }

    addBody(bodyKey, bodyValue)

    setBodyKey('')
    setBodyValue('')
  }

  return (
    <div className="flex h-auto mt-5">
      <input type="text" name="key" onChange={(e) => setBodyKey(e.target.value)} value={bodyKey}
             placeholder="Body key" className="input input-bordered"/>
      <input type="text" name="value" onChange={(e) => setBodyValue(e.target.value)} value={bodyValue}
             placeholder="Body value" className="input input-bordered ml-2.5"/>
      <button className="btn btn-accent ml-2.5" onClick={onClick}>Add</button>
    </div>
  )
}

export default Body
