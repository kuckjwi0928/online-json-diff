import React, { useState } from "react";

function QueryString(props) {
  const { addQueryString } = props
  const [queryStringKey, setQueryStringKey] = useState('');
  const [queryStringValue, setQueryStringValue] = useState('');

  const onClick = () => {
    if (!queryStringKey || !queryStringValue) {
      return
    }

    addQueryString(queryStringKey, queryStringValue)

    setQueryStringKey('')
    setQueryStringValue('')
  }

  return (
    <div className="flex h-auto mt-5">
      <input type="text" name="key" onChange={(e) => setQueryStringKey(e.target.value)} value={queryStringKey}
             placeholder="Query String Key" className="input input-bordered"/>
      <input type="text" name="value" onChange={(e) => setQueryStringValue(e.target.value)} value={queryStringValue}
             placeholder="Query String value" className="input input-bordered ml-2.5"/>
      <button className="btn btn-accent ml-2.5" onClick={onClick}>Apply</button>
    </div>
  )
}

export default QueryString
