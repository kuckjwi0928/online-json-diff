import React, { useEffect, useState } from 'react';
import storage from "../lib/LocalObjectStorage";
import ViewBadge from "./ViewBadge";
import Header from "./Header";
import Body from "./Body";
import Method from "./Method";
import ViewBody from "./ViewBody";
import QueryString from "./QueryString";

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
  const [queryStrings, setQueryStrings] = useState({});

  useEffect(() => {
    parseQueryString(originURL)
  }, [originURL]);


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

  const addQueryString = (key, value) => {
    const newObj = {
      ...queryStrings,
      [key]: value,
    }
    const queryString = toQueryString(newObj)
    setOriginURL(findDomain(originURL) + queryString)
    setCompareURL(findDomain(compareURL) + queryString)
    setQueryStrings(newObj)
  }

  const deleteQueryString = (key) => {
    delete queryStrings[key]
    const queryString = toQueryString(queryStrings)
    setOriginURL(findDomain(originURL) + queryString)
    setCompareURL(findDomain(compareURL) + queryString)
    setQueryStrings({
      ...queryStrings,
    })
  }

  const findDomain = (url) => {
    const index = url.indexOf('?')
    return url.substring(0, index === -1 ? url.length : index)
  }

  const toQueryString = (queryStringObject) => {
    if (Object.keys(queryStringObject).length <= 0) {
      return ''
    }
    return '?' + Object.entries(queryStringObject).map(([key, value]) => `${key}=${value}`).join('&')
  }

  const parseQueryString = (url) => {
    const index = url.indexOf('?')
    if (index === -1) {
      setQueryStrings({})
      return
    }
    const queries = url.substring(index + 1).split('&')
    if (queries.length > 0) {
      const obj = {}
      queries.forEach(query => {
        const [key, value] = query.split('=')
        obj[key] = value ?? ''
      })
      setQueryStrings(obj)
    }
  }

  return (
    <div className="w-full p-5">
      <Method method={method} setMethod={setMethod} />
      <Header addHeader={addHeader} />
      <ViewBadge items={headers} deleteItem={deleteHeader} />
      {
        method === 'GET' &&
          <>
            <QueryString addQueryString={addQueryString} />
            <ViewBadge items={queryStrings} deleteItem={deleteQueryString} />
          </>
      }
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
