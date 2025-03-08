import { FC, useEffect, useState } from 'react';
import { useParams, useNavigate } from 'react-router-dom';
import { Card, Typography, Spin, Button, message, Divider } from 'antd';
import ReactMarkdown from 'react-markdown';
import remarkGfm from 'remark-gfm';
import rehypeRaw from 'rehype-raw';

const { Title } = Typography;

interface Post {
  id: number;
  title: string;
  content: string;
}

const ViewPost: FC = () => {
  const { id } = useParams<{ id: string }>();
  const navigate = useNavigate();
  const [post, setPost] = useState<Post | null>(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    const fetchPost = async () => {
      try {
        setLoading(true);
        setError(null);
        
        const response = await fetch(`/api/posts/${id}`);
        
        if (!response.ok) {
          throw new Error(`Failed to fetch post (${response.status})`);
        }
        
        const data = await response.json();
        setPost(data.post);
      } catch (err) {
        console.error('Error fetching post:', err);
        setError('Failed to load post. Please try again later.');
        message.error('Failed to load post');
      } finally {
        setLoading(false);
      }
    };

    if (id) {
      fetchPost();
    }
  }, [id]);

  const handleBack = () => {
    navigate('/');
  };

  if (loading) {
    return (
      <div className="flex justify-center items-center h-64">
        <Spin size="large" tip="Loading post..." />
      </div>
    );
  }

  if (error || !post) {
    return (
      <div className="max-w-3xl mx-auto p-4">
        <Card>
          <div className="text-center">
            <Title level={4} type="danger">{error || 'Post not found'}</Title>
            <Button onClick={handleBack}>Back to Home</Button>
          </div>
        </Card>
      </div>
    );
  }

  return (
    <div className="max-w-3xl mx-auto p-4">
      <Card>
        <Title level={2}>{post.title}</Title>
        <div className="text-gray-500 mb-2">Post ID: {post.id}</div>
        <Divider />
        <div className="markdown-body">
          <ReactMarkdown 
            remarkPlugins={[remarkGfm]} 
            rehypePlugins={[rehypeRaw]}
          >
            {post.content}
          </ReactMarkdown>
        </div>
        <div className="mt-6">
          <Button type="primary" onClick={handleBack}>
            Back to Home
          </Button>
        </div>
      </Card>
    </div>
  );
};

export default ViewPost;
