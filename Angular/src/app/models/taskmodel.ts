export interface TaskModel {
  id?: number;
  created_at?: string;
  updated_at?: string;
  title?: string;
  status?: boolean;
  statusText?: string;
}

export interface CreateTask {
  id?: number;
  title?: string;
  status?: boolean;
}
