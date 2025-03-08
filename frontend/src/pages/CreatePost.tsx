import { FC, useState } from 'react';
import { Form, Input, Button, Card, message, Typography } from 'antd';
import { useNavigate } from 'react-router-dom';
import MDEditor from '@uiw/react-md-editor';

const { Title } = Typography;

interface PostFormData {
  title: string;
  content: string;
}

const CreatePost: FC = () => {
  const [loading, setLoading] = useState(false);
  const navigate = useNavigate();
  const [form] = Form.useForm();
  const [mdContent, setMdContent] = useState<string | undefined>("");

  const handleSubmit = async (values: PostFormData) => {
    if (!mdContent) {
      message.error('Please enter content');
      return;
    }

    setLoading(true);
    try {
      const response = await fetch('/api/posts', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({
          post: {
            title: values.title,
            content: mdContent
          }
        }),
      });

      if (!response.ok) {
        throw new Error('Failed to create post');
      }

      const data = await response.json();
      message.success('Post created successfully!');
      
      // Navigate to view the created post using the correct response structure
      navigate(`/posts/${data.post.id}`);
    } catch (error) {
      console.error('Error creating post:', error);
      message.error('Failed to create post. Please try again.');
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="max-w-3xl mx-auto p-4">
      <Card>
        <Title level={2}>Create New Post</Title>
        <Form
          form={form}
          layout="vertical"
          onFinish={handleSubmit}
          autoComplete="off"
        >
          <Form.Item
            name="title"
            label="Title"
            rules={[{ required: true, message: 'Please enter a title' }]}
          >
            <Input placeholder="Enter post title" />
          </Form.Item>

          <Form.Item
            label="Content"
            required
            help="Support Markdown format"
          >
            <div data-color-mode="light">
              <MDEditor
                value={mdContent}
                onChange={setMdContent}
                height={400}
                preview="edit"
              />
            </div>
          </Form.Item>

          <Form.Item>
            <Button type="primary" htmlType="submit" loading={loading} block>
              Create Post
            </Button>
          </Form.Item>
        </Form>
      </Card>
    </div>
  );
};

export default CreatePost;
