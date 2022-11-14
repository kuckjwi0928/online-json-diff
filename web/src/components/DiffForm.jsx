import React, { useState } from 'react';
import storage from "../lib/LocalObjectStorage";
import ViewBadge from "./ViewBadge";
import Header from "./Header";
import Body from "./Body";
import Method from "./Method";
import ViewBody from "./ViewBody";

function DiffForm(props) {
  const {afterClick} = props

  const latestDiff = storage.get('latest-diff') ?? {}
  const [method, setMethod] = useState(latestDiff.method ?? 'GET');
  const [headers, setHeaders] = useState({...latestDiff.headers})
  const [body, setBody] = useState(latestDiff.body)
  const [originURL, setOriginURL] = useState(latestDiff.originURL)
  const [compareURL, setCompareURL] = useState(latestDiff.compareURL)
  const [originError, setOriginError] = useState('');
  const [compareError, setCompareError] = useState('');

  const onClick = () => {
    if (!originURL) {
      setOriginError('input-error')
      return
    }

    if (!compareURL) {
      setCompareError('input-error')
      return
    }

    setOriginError('')
    setCompareError('')

    afterClick({
      originURL,
      compareURL,
      method,
      body,
      headers
    })
  }

  const onInputChange = (e) => {
    if (e.target.name === 'originURL') {
      setOriginURL(e.target.value)
      setOriginError('')
    } else {
      setCompareURL(e.target.value)
      setCompareError('')
    }
  }

  const addHeader = (key, value) => {
    setHeaders({
      ...headers,
      [key]: value,
    })
  }

  const deleteHeader = (key) => {
    delete headers[key]
    setHeaders({
      ...headers,
    })
  }

  return (
    <div className="w-full p-5">
      <Method method={method} setMethod={setMethod} />
      <Header addHeader={addHeader} />
      <ViewBadge items={headers} deleteItem={deleteHeader} />
      {
        method === 'POST' &&
          <>
            <Body body={body} setBody={setBody} />
            <ViewBody body={body} />
          </>
      }
      <div className="divider" />
      <div className="grid grid-cols-3 h-auto justify-items-center">
        <input type="text" name="originURL" onChange={onInputChange} placeholder="https://"
               className={`input input-bordered ${originError} w-full`} value={originURL}/>
        <button className="btn btn-accent" onClick={onClick}>Diff</button>
        <input type="text" name="compareURL" onChange={onInputChange} placeholder="https://"
               className={`input input-bordered ${compareError} w-full`} value={compareURL}/>
      </div>
    </div>
  )
}

export default DiffForm
