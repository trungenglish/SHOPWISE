import React, { useEffect, useState } from 'react';
import { listSessions, DecisionSession, deleteSession, renameSession } from '../api/decisionMemory';
import { Button } from './ui/button';
import { Input } from './ui/input';
import { Search, Trash, Edit2, Archive, MessageSquare } from 'lucide-react';

export function SessionHistorySidebar({ onSelectSession }: { onSelectSession: (id: string) => void }) {
  const [sessions, setSessions] = useState<DecisionSession[]>([]);
  const [search, setSearch] = useState('');
  const [editingId, setEditingId] = useState<string | null>(null);
  const [editTitle, setEditTitle] = useState('');

  useEffect(() => {
    load();
  }, []);

  const load = async () => {
    try {
      const res = await listSessions();
      setSessions(res.data || []);
    } catch (err) {
      console.error(err);
    }
  };

  const handleDelete = async (id: string) => {
    try {
      await deleteSession(id);
      load();
    } catch (err) {
      console.error(err);
    }
  };

  const handleRename = async (id: string) => {
    try {
      await renameSession(id, editTitle);
      setEditingId(null);
      load();
    } catch (err) {
      console.error(err);
    }
  };

  const filtered = sessions.filter(s => s.Title.toLowerCase().includes(search.toLowerCase()));

  // Sort by updated at (newest first)
  filtered.sort((a, b) => new Date(b.UpdatedAt).getTime() - new Date(a.UpdatedAt).getTime());

  return (
    <div className="w-64 border-r h-full flex flex-col bg-background">
      <div className="p-4 border-b">
        <h2 className="font-semibold mb-2">Decision Memory</h2>
        <div className="relative">
          <Search className="absolute left-2 top-2.5 h-4 w-4 text-muted-foreground" />
          <Input 
            placeholder="Search sessions..." 
            value={search}
            onChange={(e) => setSearch(e.target.value)}
            className="pl-8"
          />
        </div>
      </div>
      
      <div className="flex-1 overflow-y-auto p-2 space-y-1">
        {filtered.map(s => (
          <div key={s.ID} className="group flex items-center justify-between p-2 hover:bg-accent rounded-md text-sm cursor-pointer" onClick={() => onSelectSession(s.ID)}>
            {editingId === s.ID ? (
              <Input
                autoFocus
                value={editTitle}
                onChange={e => setEditTitle(e.target.value)}
                onBlur={() => handleRename(s.ID)}
                onKeyDown={e => e.key === 'Enter' && handleRename(s.ID)}
                className="h-7 text-xs"
                onClick={e => e.stopPropagation()}
              />
            ) : (
              <div className="flex items-center gap-2 overflow-hidden">
                <MessageSquare className="h-4 w-4 shrink-0 opacity-70" />
                <span className="truncate">{s.Title || 'New Session'}</span>
              </div>
            )}
            
            <div className="flex opacity-0 group-hover:opacity-100 transition-opacity" onClick={e => e.stopPropagation()}>
              <button 
                className="p-1 hover:text-blue-500"
                onClick={() => {
                  setEditingId(s.ID);
                  setEditTitle(s.Title);
                }}
              >
                <Edit2 className="h-3 w-3" />
              </button>
              <button 
                className="p-1 hover:text-red-500"
                onClick={() => handleDelete(s.ID)}
              >
                <Trash className="h-3 w-3" />
              </button>
            </div>
          </div>
        ))}
        {filtered.length === 0 && (
          <div className="p-4 text-center text-sm text-muted-foreground">
            No sessions found
          </div>
        )}
      </div>
    </div>
  );
}
