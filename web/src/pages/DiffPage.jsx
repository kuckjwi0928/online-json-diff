import React, { useState } from "react";
import DiffForm from "../components/DiffForm";
import { getV1DiffTarget } from "../http/diffApi";
import DiffViewer from "../components/DiffViewer";
import BottomNav from "../components/BottomNav";
import storage from "../lib/LocalObjectStorage";

function DiffPage() {
  const [left, setLeft] = useState('');
  const [right, setRight] = useState('');

  const diff = async ({originURL, compareURL, headers}) => {
    const {data} = await getV1DiffTarget(originURL, compareURL, headers)
    setLeft(data.left)
    setRight(data.right)

    storage.set('latest-diff', {
      originURL,
      compareURL,
      headers
    })
  }

  return (
    <>
      <DiffForm afterClick={diff}/>
      <DiffViewer left={left} right={right}/>
      <BottomNav/>
    </>
  )
}

export default DiffPage;
