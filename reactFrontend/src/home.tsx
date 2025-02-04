import UserForm from "./componets/userForm";
import UserList from "./componets/userList";

export default function Home() {
  return (
    <>
      <div className="">
        Hello
        <UserList />
        <UserForm />
      </div>
    </>
  );
}
