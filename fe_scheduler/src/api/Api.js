import axios from "axios";

const tasks = {
  /** api for scheduling a task */
  schedule: body =>
    axios.post("/api/v1/task/schedule", {
      task: {
        name: "Schedule Messages",
        schedule_at: getScheduledTime(),
        processes: [
          {
            name: "Message 1",
            work_done: 3
          },
          {
            name: "Message 2",
            work_done: 5
          }
        ]
      }
    }),

  /** api for running a task */
  sendNow: body =>
    axios.post("/api/v1/task/schedule", {
      task: {
        name: "Send Messages",
        schedule_at: "2019-10-10T10:41:02.335Z",
        processes: [
          {
            name: "Message 1",
            work_done: 2
          },
          {
            name: "Message 2",
            work_done: 7
          },
          {
            name: "Message 3",
            work_done: 1
          }
        ]
      }
    })
};
function getScheduledTime() {
  var t = new Date();
  t.setSeconds(t.getSeconds() + 10);
  return t;
}

export default {
  tasks
};
