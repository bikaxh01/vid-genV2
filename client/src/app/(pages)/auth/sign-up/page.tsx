"use client";

import { registerUser } from "@/handler/user-apis";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { EmailValidationMessage, ValidateEmail } from "@/utils/utils";

import Image from "next/image";
import Link from "next/link";
import { useRouter } from "next/navigation";
import React, { useEffect, useState } from "react";
import { toast } from "sonner";
import { useDebounceCallback } from "usehooks-ts";

function SignUp() {
  const [debouncedEmail, setDebounceEmail] = useState("");
  const [userName, setUserName] = useState("");
  const [password, setPassword] = useState("");
  const [confirmPassword, setConfirmPassword] = useState("");
  const [disableSubmit, setDisableSubmit] = useState(true);
  const [passwordError, setPasswordError] = useState("");
  const [emailValidationMessage, setEmailValidationMessage] =
    useState<EmailValidationMessage | null>(null);
  const debounced = useDebounceCallback(setDebounceEmail, 500);
  const [email, setEmail] = useState("");
  const router = useRouter();
  useEffect(() => {
    if (confirmPassword && password) {
      if (confirmPassword !== password) {
        setPasswordError("confirm password do not match");
      } else {
        setPasswordError("");
      }
    }
  }, [confirmPassword]);

  useEffect(() => {
    if (!emailValidationMessage?.success || passwordError) {
      setDisableSubmit(true);
    } else if (debouncedEmail && userName && password && confirmPassword) {
      setDisableSubmit(false);
    } else {
      setDisableSubmit(true);
    }
  }, [
    emailValidationMessage,
    passwordError,
    confirmPassword,
    password,
    userName,
    debouncedEmail,
  ]);

  useEffect(() => {
    if (debouncedEmail) {
      const checkEmail = async () => {
        const res = await ValidateEmail(debouncedEmail);
        setEmailValidationMessage(res);
      };
      checkEmail();
    }
  }, [debouncedEmail]);

  const handleSubmit = async (e: React.SyntheticEvent<HTMLFormElement>) => {
    e.preventDefault();

    try {
      setDisableSubmit(true);
      const res = await registerUser(debouncedEmail, userName, password);
      toast.success(res.message);
      router.push("/auth/sign-in");
    } catch (error: any) {
      toast.error(error);
    } finally {
      setDisableSubmit(false);
    }
  };
  return (
    <div className=" ">
      <div className=" items-center justify-center h-fit w-fit mt-4 ml-3 flex   gap-2 ">
        <Image
          alt="logo"
          src={"/logo.webp"}
          className=" size-8"
          height={100}
          width={100}
        />
        <span className="  font-semibold">ManiMate</span>
      </div>

      <div className=" h-full  w-full flex items-center justify-center ">
        <div className=" h-full w-[24rem]   px-2">
          <div className=" flex flex-col gap-4  items-center justify-center">
            <div>
              <Image
                alt="logo"
                src={"/logo.webp"}
                className=" size-20"
                height={400}
                width={400}
              />
            </div>
            <div className="   gap-2 flex flex-col items-center justify-center">
              <h1 className=" font-bold text-2xl">Sign up to ManiMate</h1>
              <div className=" flex flex-col  items-center justify-center text-xs text-neutral-400">
                <p>Get 1000 credits + 300 daily credits</p>
                <p>No waitlist—start creating now</p>
              </div>
            </div>
          </div>
          <div className=" flex flex-col gap-3 w-full mt-4 ">
            <div>
              <Button className=" bg-sidebar-accent hover:bg-sidebar-accent/70 w-full">
                <span>
                  <Image
                    src={"/google.svg"}
                    width={100}
                    height={100}
                    className=" size-4"
                    alt="google"
                  />
                </span>
                Sign up with Google
              </Button>
            </div>
            <div>
              <Button className=" bg-sidebar-accent hover:bg-sidebar-accent/70 w-full">
                <span>
                  <Image
                    src={"/github.png"}
                    width={100}
                    height={100}
                    className=" size-4"
                    alt="google"
                  />
                </span>
                Sign up with Github
              </Button>
            </div>
          </div>

          <div className="flex items-center gap-2 my-4 w-full">
            <div className="flex-1 h-px bg-neutral-800" />
            <span className="text-neutral-600 text-sm font-medium px-2">
              or
            </span>
            <div className="flex-1 h-px bg-neutral-800" />
          </div>

          <form onSubmit={handleSubmit} className="  flex flex-col gap-4">
            <div className=" flex flex-col gap-2">
              <label className=" text-xs">Email</label>
              <Input
                className=" rounded-lg "
                value={email}
                onChange={(e) => {
                  setEmail(e.target.value);
                  debounced(e.target.value);
                }}
              />
              {emailValidationMessage && (
                <span
                  className={` text-xs ${
                    emailValidationMessage.success
                      ? "text-green-500"
                      : "text-red-500"
                  } `}
                >
                  {emailValidationMessage.message}
                </span>
              )}
            </div>
            <div className=" flex flex-col gap-2">
              <label className=" text-xs">Username</label>
              <Input
                className=" rounded-lg"
                value={userName}
                onChange={(e) => setUserName(e.target.value)}
              />
            </div>
            <div className=" flex flex-col gap-2">
              <label className=" text-xs">Password</label>
              <Input
                className=" rounded-lg"
                value={password}
                onChange={(e) => setPassword(e.target.value)}
              />
            </div>
            <div className=" flex flex-col gap-2">
              <label className=" text-xs">Confirm Password</label>
              <Input
                className=" rounded-lg"
                value={confirmPassword}
                onChange={(e) => setConfirmPassword(e.target.value)}
              />
              {passwordError && (
                <span className=" text-xs  text-red-500">{passwordError}</span>
              )}
            </div>

            <Button type="submit" disabled={disableSubmit}>
              Sign Up
            </Button>
          </form>
          <div className=" text-xs text-neutral-500  flex  items-center justify-center mt-5 gap-2">
            Already have an account?{" "}
            <Link href={"/auth/sign-in"} className=" underline">
              Sign-in
            </Link>
          </div>
        </div>
      </div>
    </div>
  );
}

export default SignUp;
