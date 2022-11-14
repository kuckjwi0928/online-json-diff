import React, { useState } from "react";
import DiffForm from "../components/DiffForm";
import { getV1DiffTarget } from "../http/diffApi";
import DiffViewer from "../components/DiffViewer";
import BottomNav from "../components/BottomNav";
import storage from "../lib/LocalObjectStorage";
import LoadingBar from "../components/LoadingBar";

function DiffPage() {
  const [left, setLeft] = useState('');
  const [right, setRight] = useState('');
  const [show, setShow] = useState(false)

  const diff = async ({originURL, compareURL, method, body, headers}) => {
    setShow(true)

    const {data} = await getV1DiffTarget(originURL, compareURL, method, body, headers)

    setLeft(data.left)
    setRight(data.right)

    storage.set('latest-diff', {
      originURL,
      compareURL,
      method,
      body,
      headers
    })

    setShow(false)
  }

  return (
    <>
      <DiffForm afterClick={diff}/>
      <DiffViewer left={left} right={right}/>
      <BottomNav/>
      <LoadingBar show={show} />
    </>
  )
}

export default DiffPage;
