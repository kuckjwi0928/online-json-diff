import React, { useState } from 'react';
import storage from "../lib/LocalObjectStorage";

function DiffForm(props) {
  const {afterClick} = props

  const latestDiff = storage.get('latest-diff') ?? {}
  const [headers, setHeaders] = useState({...latestDiff.headers})
  const [originURL, setOriginURL] = useState(latestDiff.originURL)
  const [compareURL, setCompareURL] = useState(latestDiff.compareURL)
  const [originError, setOriginError] = useState('');
  const [compareError, setCompareError] = useState('');
  const [headerKey, setHeaderKey] = useState('');
  const [headerValue, setHeaderValue] = useState('');

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

  const addHeader = () => {
    if (!headerKey || !headerValue) {
      return
    }
    setHeaders({
      ...headers,
      [headerKey]: headerValue,
    })
    setHeaderKey('')
    setHeaderValue('')
  }

  const deleteHeader = (key) => {
    delete headers[key]
    setHeaders({
      ...headers,
    })
  }

  return (
    <div className="w-full p-5">
      <div className="flex h-auto">
        <input type="text" name="key" onChange={(e) => setHeaderKey(e.target.value)} value={headerKey}
               placeholder="Header key" className="input input-bordered"/>
        <input type="text" name="value" onChange={(e) => setHeaderValue(e.target.value)} value={headerValue}
               placeholder="Header value" className="input input-bordered ml-2.5"/>
        <button className="btn btn-accent ml-2.5" onClick={addHeader}>Add</button>
      </div>
      <div className="mt-5">
        {
          Object.keys(headers).map((key, index) =>
            <button className={`badge badge-outline ${index !== 0 ? 'ml-1.5' : ''}`} key={key} onClick={() => deleteHeader(key)}>
              {`${key}=${headers[key]}`}
            </button>
          )
        }
      </div>
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
