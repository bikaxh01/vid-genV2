"use client"
import React, { useState } from "react";
import {
  Accordion,
  AccordionContent,
  AccordionItem,
  AccordionTrigger,
} from "@/components/ui/accordion";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";

function SceneDescriptionCard({
  title,
  description,
  sequence,
}: {
  title: string;
  description: string;
  sequence: number;
}) {

  const [input,setInput] = useState("title")
  return (
    <div>
      <Accordion type="single" collapsible>
        <AccordionItem value={sequence.toString()}>
          <AccordionTrigger>
            <div className=" flex  w-full justify-between">
              <h2>{title}</h2>
              <span className="  text-xs text-neutral-500">
                scene {sequence}
              </span>
            </div>
          </AccordionTrigger>
          <AccordionContent className=" flex flex-col gap-2 ">
            <div className=" flex flex-col gap-1">
              <Label>Title</Label>
              <Input value={input} onChange={(e)=>setInput(e.target.value)} className="!text-xs" />
            </div>
            <div className=" flex-col flex gap-1">
              <Label>Description</Label>
              <Input value={description} className=" !text-xs" />
            </div>
          </AccordionContent>
        </AccordionItem>
      </Accordion>
    </div>
  );
}

export default SceneDescriptionCard;
