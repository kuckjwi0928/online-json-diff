import React, { useState } from "react";

function Header(props) {
  const { addHeader } = props
  const [headerKey, setHeaderKey] = useState('');
  const [headerValue, setHeaderValue] = useState('');

  const onClick = () => {
    if (!headerKey || !headerValue) {
      return
    }

    addHeader(headerKey, headerValue)

    setHeaderKey('')
    setHeaderValue('')
  }

  return (
    <div className="flex h-auto mt-5">
      <input type="text" name="key" onChange={(e) => setHeaderKey(e.target.value)} value={headerKey}
             placeholder="Header key" className="input input-bordered"/>
      <input type="text" name="value" onChange={(e) => setHeaderValue(e.target.value)} value={headerValue}
             placeholder="Header value" className="input input-bordered ml-2.5"/>
      <button className="btn btn-accent ml-2.5" onClick={onClick}>Add</button>
    </div>
  )
}

export default Header
